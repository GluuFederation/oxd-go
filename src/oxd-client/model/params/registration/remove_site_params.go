//  Author: Michał Kępkowski
//  Date: 07.04.2018

package model

// Remove site request https://gluu.org/docs/oxd/3.1.2/api/#remove-site
type RemoveSiteRequestParams struct {
	OxdId string `json:"oxd_id"`
}

// Remove site response https://gluu.org/docs/oxd/3.1.2/api/#remove-site
type RemoveSiteResponseParams struct {
	OxdId string `json:"oxd_id"`
}