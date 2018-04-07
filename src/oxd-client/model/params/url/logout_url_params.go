//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package model

// Get logout url request https://gluu.org/docs/oxd/3.1.2/api/#get-logout-uri
type LogoutUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	IdTokenHint string `json:"id_token_hint,omitempty"`
	PostLogoutRedirectUri string `json:"post_logout_redirect_uri,omitempty"`
	State string `json:"state,omitempty"`
	SessionState string `json:"session_state,omitempty"`
}

// Get logout url response https://gluu.org/docs/oxd/3.1.2/api/#get-logout-uri
type LogoutUrlResponseParams struct {
	Uri string `json:"uri"`

}
