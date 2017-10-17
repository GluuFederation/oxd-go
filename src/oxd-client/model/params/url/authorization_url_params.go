//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

type AuthorizationUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	AcrValues []string  `json:"acr_values"`
	Scope string `json:"scope,omitempty"`
	
	Prompt string `json:"prompt"`
	CustomParameters 
	
}
type CustomParameters map[string]string

type AuthorizationUrlResponseParams struct {
	AuthorizationUrl string `json:"authorization_url"`
}
