//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package model

// Get user info request https://gluu.org/docs/oxd/3.1.2/api/#get-user-info
type UserInfoRequestParams struct {
	OxdId       string `json:"oxd_id"`
	ProtectionAccessToken string `json:"protection_access_token"`
	AccessToken string `json:"access_token"`
}

// Get user info response https://gluu.org/docs/oxd/3.1.2/api/#get-user-info
type UserInfoResponseParams struct {
	Claims map[string] []string `json:"claims"`
}
