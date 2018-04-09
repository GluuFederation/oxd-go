//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package uma

import "oxd-client/model/params/uma/protect"

// Protect resource request https://gluu.org/docs/oxd/3.1.2/api/#uma-rs-protect-resources
type RsProtectRequestParams struct {
	OxdId string `json:"oxd_id"`
	Resources []protect.RsResource `json:"resources"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

// Protect resource response https://gluu.org/docs/oxd/3.1.2/api/#uma-rs-protect-resources
type RsProtectResponseParams struct {
	Status string `json:"status"`
	OxdId string `json:"oxd_id"`
}
