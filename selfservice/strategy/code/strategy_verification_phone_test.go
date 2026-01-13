package code_test

import (
	"context"
	_ "embed"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/x/configx"

	"github.com/ory/kratos/driver/config"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/internal/testhelpers"
	"github.com/ory/kratos/selfservice/flow/verification"
	"github.com/ory/kratos/text"
	"github.com/ory/kratos/ui/node"
	"github.com/ory/kratos/x"
	"github.com/ory/x/ioutilx"
	"github.com/ory/x/sqlxx"
)

func TestVerificationPhone(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	conf, reg := internal.NewFastRegistryWithMocks(t, configx.WithValues(defaultConfig))

	verificationPhone := "+1234567890"

	identityToVerify := &identity.Identity{
		ID:       x.NewUUID(),
		Traits:   identity.Traits(`{"email":"verifyme@ory.sh", "phone":"` + verificationPhone + `"}`),
		SchemaID: config.DefaultIdentityTraitsSchemaID,
		Credentials: map[identity.CredentialsType]identity.Credentials{
			"password": {
				Type:        "password",
				Identifiers: []string{"recoverme@ory.sh"},
				Config:      sqlxx.JSONRawMessage(`{"hashed_password":"foo"}`),
			},
		},
	}

	_ = testhelpers.NewVerificationUIFlowEchoServer(t, reg)
	_ = testhelpers.NewLoginUIFlowEchoServer(t, reg)
	_ = testhelpers.NewSettingsUIFlowEchoServer(t, reg)
	_ = testhelpers.NewErrorTestServer(t, reg)
	_ = testhelpers.NewRedirTS(t, "returned", conf)

	public, _ := testhelpers.NewKratosServerWithCSRF(t, reg)

	require.NoError(t, reg.IdentityManager().Create(context.Background(), identityToVerify,
		identity.ManagerAllowWriteProtectedTraits))

	expectSuccess := func(t *testing.T, hc *http.Client, isAPI, isSPA bool, values func(url.Values)) string {
		if hc == nil {
			hc = testhelpers.NewDebugClient(t)
			if !isAPI {
				hc = testhelpers.NewClientWithCookies(t)
				hc.Transport = testhelpers.NewTransportWithLogger(http.DefaultTransport, t).RoundTripper
			}
		}

		return testhelpers.SubmitVerificationForm(t, isAPI, isSPA, hc, public, values, http.StatusOK,
			testhelpers.ExpectURL(isAPI || isSPA, public.URL+verification.RouteSubmitFlow, conf.SelfServiceFlowVerificationUI(ctx).String()))
	}

	submitVerificationCode := func(t *testing.T, body string, c *http.Client, code string) (string, *http.Response) {
		action := gjson.Get(body, "ui.action").String()
		require.NotEmpty(t, action, "%v", string(body))
		csrfToken := extractCsrfToken([]byte(body))

		res, err := c.PostForm(action, url.Values{
			"code":       {code},
			"csrf_token": {csrfToken},
		})
		require.NoError(t, err)

		return string(ioutilx.MustReadAll(res.Body)), res
	}

	t.Run("description=should verify a phone number", func(t *testing.T) {
		check := func(t *testing.T, actual string) {
			assert.EqualValues(t, string(node.CodeGroup), gjson.Get(actual, "active").String(), "%s", actual)
			assert.EqualValues(t, verificationPhone, gjson.Get(actual, "ui.nodes.#(attributes.name==phone).attributes.value").String(), "%s", actual)
			// assertx.EqualAsJSON(t, text.NewVerificationEmailWithCodeSent(), json.RawMessage(gjson.Get(actual, "ui.messages.0").Raw))

			// Verify that an SMS message was "sent" (queued in courier)
			message := testhelpers.CourierExpectMessage(ctx, t, reg, verificationPhone, "Your verification code is")
			// SMS templates usually include the code directly
			code := testhelpers.CourierExpectCodeInMessage(t, message, 1)

			cl := testhelpers.NewClientWithCookies(t)

			// We submit the code manually since SMS doesn't have a magic link usually
			body, res := submitVerificationCode(t, actual, cl, code)

			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.EqualValues(t, "passed_challenge", gjson.Get(body, "state").String())
			assert.EqualValues(t, text.NewInfoSelfServiceVerificationSuccessful().Text, gjson.Get(body, "ui.messages.0.text").String())

			id, err := reg.PrivilegedIdentityPool().GetIdentityConfidential(context.Background(), identityToVerify.ID)
			require.NoError(t, err)

			// Find the verified address
			var found bool
			for _, address := range id.VerifiableAddresses {
				if address.Value == verificationPhone {
					assert.True(t, address.Verified)
					assert.EqualValues(t, identity.VerifiableAddressStatusCompleted, address.Status)
					found = true
					break
				}
			}
			assert.True(t, found, "Verifiable address not found")
		}

		values := func(v url.Values) {
			v.Set("phone", verificationPhone)
		}

		t.Run("type=browser", func(t *testing.T) {
			check(t, expectSuccess(t, nil, false, false, values))
		})
	})
}
