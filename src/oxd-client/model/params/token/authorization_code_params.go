//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

// Deprecated:
type AuthorizationCodeRequestParams struct {
	OxdId string `json:"oxd_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	AcrValues []string `json:"acr_values,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	State string `json:"state"`
    
}
// Deprecated:
type AuthorizationCodeResponseParams struct {
	Code string `json:"code"`
}
