//  Author: Michał Kępkowski
//  Date: 07.04.2018

package uma

// Introspect Rpt request https://gluu.org/docs/oxd/3.1.2/api/#uma-2-introspect-rpt
type IntrospectRptRequestParams struct {
	OxdId string `json:"oxd_id"`
	Rpt string `json:"rpt"`
}

// Introspect Rpt response https://gluu.org/docs/oxd/3.1.2/api/#uma-2-introspect-rpt
type IntrospectRptResponseParams struct {
	Active bool `json:"active"`
	Exp uint `json:"exp"`
	Iat uint `json:"iat"`
	Permissions []Permission `json:"permissions"`
}

// Permissions structure used in IntrospectRptResponseParams
type Permission struct {
	ResourceId string `json:"resource_id"`
	ResourceScopes []string `json:"resource_scopes"`
	Exp uint `json:"exp"`
}