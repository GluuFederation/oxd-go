package model

// import "encoding/json"

type GetClientTokenRequestParams struct {
	OxdId string `json:"oxd_id"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OpHost string `json:"op_host"`
	OpDiscoveryPath string `json:"op_discovery_path,omitempty"`
	Scope string `json:"scope,omitempty"`
	
}


type GetClientTokenResponseParams struct {
	
	Scope string `json:"scope"`
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
   
}


