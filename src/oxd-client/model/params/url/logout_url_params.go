//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

type LogoutUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	IdTokenHint string `json:"id_token_hint,omitempty"`
	PostLogoutRedirectUri string `json:"post_logout_redirect_uri,omitempty"`
	State string `json:"state,omitempty"`
	SessionState string `json:"session_state,omitempty"`
}

type LogoutUrlResponseParams struct {
	Uri string `json:"uri"`

}
