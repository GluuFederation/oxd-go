//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package uma

type RpGetRptRequestParams struct {
	OxdId string `json:"oxd_id"`
	Ticket string `json:"ticket"`
	ClaimToken string `json:"claim_token,omitempty"`
	ClaimTokenFormat string `json:"claim_token_format,omitempty"`
	Pct string `json:"pct,omitempty"`
	Rpt string `json:"access_token,omitempty"`
	Scope string `json:"scope,omitempty"`
	State string `json:"state,omitempty"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

type RpGetRptResponseParams struct {
	Rpt string `json:"access_token"`
	TokenType string `json:"token_type"`
	Pct string `json:"pct"`
	Upgraded bool `json:"updated"`
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

