//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

type UserInfoRequestParams struct {
	OxdId       string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	AccessToken string `json:"access_token"`
}

type UserInfoResponseParams struct {
	Claims map[string] []string `json:"claims"`
}
