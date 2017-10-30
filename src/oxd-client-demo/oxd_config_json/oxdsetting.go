package newmodel

type OxdSetting struct {
    OxdHost string `json:"oxd_host"`
	OpHost string  `json:"op_host"`
	OxdPort string  `json:"oxd_host_port"`
	AuthorizationRedirectUri string  `json:"authorization_redirect_uri"`
	PostLogoutRedirectUri string  `json:"post_logout_redirect_uri"`
    
	ResponseTypes []string  `json:"response_types"`

	ClientId string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	ClientName string  `json:"client_name"`
	ConnectionType string  `json:"connection_type"`
	ApplicationType string  `json:"application_type"`
	Scope  []string  `json:"scope"`
	
	AcrValues  []string  `json:"acr_values"`
	GrantType  []string  `json:"grant_types"`
	
	OxdId string `json:"oxd_id"`
	DynamicRegistration string `json:"dynamic_registration"` 
	HttpRestUrl string `json:"https_rest_url"`
}

