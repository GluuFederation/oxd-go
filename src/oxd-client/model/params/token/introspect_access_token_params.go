//  Author: Michał Kępkowski
//  Date: 07.04.2018

package model

// Introspect Access Token request https://gluu.org/docs/oxd/3.1.2/api/#introspect-access-token
type IntrospectAccessTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	AccessToken string `json:"access_token"`
}

// Introspect Access Token response https://gluu.org/docs/oxd/3.1.2/api/#introspect-access-token
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