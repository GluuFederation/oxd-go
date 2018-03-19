package model

type IntrospectAccessTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	AccessToken string `json:"access_token"`
}


type IntrospectAccessTokenResponseParams struct {
	Active bool `json:"active"`
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Scopes []string `json:"scopes"`
	TokenType string `json:"token_type"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iss string `json:"iss"`
	Exp int `json:"exp"`
	Iat int `json:"iat"`
	AcrValues []string `json:"acr_values"`
	ExtensionField []string `json:"extension_field"`
}