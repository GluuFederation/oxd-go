//  Author: Michał Kępkowski
//  Date: 07.04.2018

package model

// Get Access Token by Refresh Token request https://gluu.org/docs/oxd/3.1.2/api/#get-access-token-by-refresh-token
type GetAccessTokenByRefreshTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	RefreshToken	string `json:"refresh_token"`
	ProtectionAccessToken string `json:"protection_access_token"`
	Scope []string `json:"scope,omitempty"`
}

// Get Access Token by Refresh Token response https://gluu.org/docs/oxd/3.1.2/api/#get-access-token-by-refresh-token
type GetAccessTokenByRefreshTokenResponseParams struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken	string `json:"refresh_token"`
}