//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package model

// Get authorization url request https://gluu.org/docs/oxd/3.1.2/api/#get-authorization-url
type AuthorizationUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	AcrValues []string  `json:"acr_values"`
	Scope []string `json:"scope"`
	
	Prompt string `json:"prompt"`
	CustomParameters
}

// Custom parameters map
type CustomParameters map[string]string

// Get authorization url response https://gluu.org/docs/oxd/3.1.2/api/#get-authorization-url
type AuthorizationUrlResponseParams struct {
	AuthorizationUrl string `json:"authorization_url"`
}
