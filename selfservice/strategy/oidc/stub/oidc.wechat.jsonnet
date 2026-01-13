// 微信 OIDC JSONNet Mapper
// 用于将微信用户信息映射到 Kratos Identity 的 traits
//
// WeChat 返回的 Claims 字段说明:
// - subject: UnionID (如已绑定开放平台) 或 OpenID
// - nickname: 用户昵称
// - picture: 头像 URL
// - gender: male/female/空
// - raw_claims.openid: 用户在当前应用的唯一标识
// - raw_claims.unionid: 用户在开放平台下的唯一标识 (跨应用一致)
// - raw_claims.province: 省份
// - raw_claims.city: 城市
// - raw_claims.country: 国家

local claims = std.extVar('claims');

{
  identity: {
    traits: {
      // 使用 UnionID 作为唯一标识 (如果没有则使用 OpenID)
      [if 'raw_claims' in claims && 'unionid' in claims.raw_claims && claims.raw_claims.unionid != '' then 'wechat_unionid']: claims.raw_claims.unionid,
      [if 'raw_claims' in claims && 'openid' in claims.raw_claims then 'wechat_openid']: claims.raw_claims.openid,
      
      // 用户名使用微信昵称
      [if 'nickname' in claims && claims.nickname != '' then 'username']: claims.nickname,
      
      // 头像
      [if 'picture' in claims && claims.picture != '' then 'avatar']: claims.picture,
    },
    
    // 可选: 设置验证地址 (微信不提供邮箱)
    // verifiable_addresses: [],
  },
}
