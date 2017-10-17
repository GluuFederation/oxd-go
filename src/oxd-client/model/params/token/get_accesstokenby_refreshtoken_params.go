package model

type GetAccessTokenByRefreshTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	RefreshToken	string `json:"refresh_token"`
	ProtectionAccessToken string `json:"protection_access_token"`
	// Scope string `json:"scope,omitempty"`
    // Code  string `json:"code,omitempty"`
	// State string `json:"state,omitempty"`
}

type GetAccessTokenByRefreshTokenResponseParams struct {
	
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	
	RefreshToken	string `json:"refresh_token"`

}