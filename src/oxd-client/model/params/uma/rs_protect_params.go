//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package uma

import "oxd-client/model/params/uma/protect"

type RsProtectRequestParams struct {
	OxdId string `json:"oxd_id"`
	Resources []protect.RsResource `json:"resources"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

type RsProtectResponseParams struct {

	Status string `json:"status"`
	OxdId string `json:"oxd_id"`
}
