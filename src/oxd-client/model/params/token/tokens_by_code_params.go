//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

type TokensByCodeRequestParams struct {
	OxdId string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	Code  string `json:"code"`
	State string `json:"state,omitempty"`
}

type TokensByCodeResponseParams struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	IdToken	string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	IdTokenClaims IdTokenClaims `json:"id_token_claims"`
}
type IdTokenClaims struct {
	Iss []string `json:"iss"`
	Sub []string `json:"sub"`
	Aud []string `json:"aud"`
	Nonce []string `json:"nonce"`
	Exp []string `json:"exp"`
	Iat []string `json:"iat"`
	AuthTime []string `json:"auth_time"`
	AtHash []string `json:"at_hash"`
	oxOpenIdCOnnectVersion []string `json:"oxOpenIDConnectVersion"`
}