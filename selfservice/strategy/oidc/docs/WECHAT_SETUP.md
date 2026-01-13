# 微信 OIDC Provider 配置指南
# WeChat Open Platform OAuth2.0 Integration for Ory Kratos

## 前置要求

1. 在微信开放平台 (https://open.weixin.qq.com/) 注册开发者账号
2. 创建"网站应用"并完成审核
3. 获取 AppID 和 AppSecret
4. 在应用设置中配置授权回调域名

## 回调 URL 配置

在微信开放平台设置授权回调域名时，需要配置：
- 回调域名: account.17cid.bei (不含协议前缀)
- 完整回调 URL: https://account.17cid.bei/self-service/methods/oidc/callback/wechat

## Kratos 配置示例

```yaml
selfservice:
  methods:
    oidc:
      enabled: true
      config:
        base_redirect_uri: https://account.17cid.bei/
        providers:
          # 微信网站应用 (PC 二维码扫码登录)
          - id: wechat
            provider: wechat
            client_id: YOUR_WECHAT_APP_ID       # 替换为微信 AppID
            client_secret: YOUR_WECHAT_APP_SECRET  # 替换为微信 AppSecret
            mapper_url: file:///etc/config/kratos/wechat-mapper.jsonnet
            scope:
              - snsapi_login  # 网站应用必须使用此 scope
```

## JSONNet Mapper 配置

将以下内容保存为 `/etc/config/kratos/wechat-mapper.jsonnet`:

```jsonnet
// 微信 OIDC JSONNet Mapper
local claims = std.extVar('claims');

{
  identity: {
    traits: {
      // UnionID (跨应用统一标识，如已绑定开放平台)
      [if 'raw_claims' in claims && 'unionid' in claims.raw_claims && claims.raw_claims.unionid != '' then 'wechat_unionid']: claims.raw_claims.unionid,
      
      // OpenID (当前应用内的用户标识)
      [if 'raw_claims' in claims && 'openid' in claims.raw_claims then 'wechat_openid']: claims.raw_claims.openid,
      
      // 用户名使用微信昵称
      [if 'nickname' in claims && claims.nickname != '' then 'username']: claims.nickname,
      
      // 头像
      [if 'picture' in claims && claims.picture != '' then 'avatar']: claims.picture,
    },
  },
}
```

## Identity Schema 更新

需要在 Identity Schema 中添加微信相关字段:

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "traits": {
      "type": "object",
      "properties": {
        "wechat_unionid": {
          "type": "string",
          "title": "微信 UnionID"
        },
        "wechat_openid": {
          "type": "string",
          "title": "微信 OpenID"
        },
        "username": {
          "type": "string",
          "title": "用户名"
        },
        "avatar": {
          "type": "string",
          "format": "uri",
          "title": "头像"
        }
      }
    }
  }
}
```

## 微信 OAuth2.0 流程说明

1. **授权请求**: 用户访问登录页面，点击"微信登录"按钮
2. **跳转二维码**: 重定向到 `https://open.weixin.qq.com/connect/qrconnect`
3. **用户扫码**: 用户使用微信 APP 扫描二维码并确认授权
4. **回调**: 微信回调到 Kratos 的 callback URL，携带 authorization code
5. **换取 Token**: Kratos 用 code 换取 access_token (同时获得 openid/unionid)
6. **获取用户信息**: 使用 access_token + openid 调用 userinfo 接口
7. **创建/关联身份**: 根据 mapper 配置创建或关联 Kratos Identity

## 返回的用户信息字段

| 字段 | 说明 |
|------|------|
| `openid` | 用户在当前应用的唯一标识 |
| `unionid` | 用户在开放平台下的唯一标识 (需绑定开放平台) |
| `nickname` | 用户昵称 |
| `headimgurl` | 用户头像 URL |
| `sex` | 性别 (1=男, 2=女, 0=未知) |
| `province` | 省份 |
| `city` | 城市 |
| `country` | 国家 |

## 注意事项

1. **UnionID**: 只有将应用绑定到微信开放平台账号后，才能获取 UnionID
2. **Scope**: 网站应用必须使用 `snsapi_login`，不能使用 `snsapi_userinfo`
3. **头像**: 微信头像 URL 可能会过期，建议下载保存到自己的存储
4. **安全**: AppSecret 必须保密，只能在服务器端使用
5. **审核**: 网站应用需要通过微信审核才能正常使用
