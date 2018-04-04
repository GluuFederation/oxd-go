//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package validation
// Deprecated:
type CheckIdTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	IdToken string `json:"id_token"`
}
// Deprecated:
type CheckIdTokenResponseParams struct {
	Active bool `json:"active"`
	ExpiresAt int64 `json:"expires_at"`
	IssuedAt int64 `json:"issued_at"`
	Claims map[string] []string `json:"claims"`
}
