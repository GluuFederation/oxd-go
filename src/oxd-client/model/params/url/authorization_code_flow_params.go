//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package model
// Deprecated:
type AuthorizationCodeFlowRequestParams struct {
	OxdId string `json:"oxd_id"`
	RedirectUrl string `json:"redirect_url,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	UserId string `json:"user_id,omitempty"`
	UserSecret string `json:"user_secret,omitempty"`
	Scope string `json:"scope,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	Acr string `json:"acr,omitempty"`
}
// Deprecated:
type AuthorizationCodeFlowResponseParams struct {
	AccessToken string `json:"access_token"`
	ExpiresInSeconds int64 `json:"expires_in_seconds"`
	RefreshToken string `json:"refresh_token"`
	AuthorizationCode string `json:"authorization_code"`
	Scope string `json:"scope"`
	IdToken string `json:"id_token"`
}

