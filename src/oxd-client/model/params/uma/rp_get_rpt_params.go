//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package uma

// Get Rpt request https://gluu.org/docs/oxd/3.1.2/api/#uma-rp-get-rpt
type RpGetRptRequestParams struct {
	OxdId                 string   `json:"oxd_id"`
	Ticket                string   `json:"ticket"`
	ClaimToken            string   `json:"claim_token,omitempty"`
	ClaimTokenFormat      string   `json:"claim_token_format,omitempty"`
	Pct                   string   `json:"pct,omitempty"`
	Rpt                   string   `json:"access_token,omitempty"`
	Scope                 []string `json:"scope,omitempty"`
	State                 string   `json:"state,omitempty"`
	ProtectionAccessToken string   `json:"protection_access_token"`
}

// Get Rpt response https://gluu.org/docs/oxd/3.1.2/api/#uma-rp-get-rpt
type RpGetRptResponseParams struct {
	Rpt              string       `json:"access_token"`
	TokenType        string       `json:"token_type"`
	Pct              string       `json:"pct"`
	Upgraded         bool         `json:"updated"`
	Error            string       `json:"error"`
	ErrorDescription string       `json:"error_description"`
	Details          ErrorDetails `json:"details"`
}

type ErrorDetails interface{}
