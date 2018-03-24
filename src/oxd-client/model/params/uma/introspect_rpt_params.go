package uma

type IntrospectRptRequestParams struct {
	OxdId string `json:"oxd_id"`
	Rpt string `json:"rpt"`
}

type IntrospectRptResponseParams struct {
	Active bool `json:"active"`
	Exp uint `json:"exp"`
	Iat uint `json:"iat"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	ResourceId string `json:"resource_id"`
	ResourceScopes []string `json:"resource_scopes"`
	Exp uint `json:"exp"`
}