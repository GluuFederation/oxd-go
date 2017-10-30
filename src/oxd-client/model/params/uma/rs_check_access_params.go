//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 03/01/17
//
package uma


type RsCheckAccessRequestParams struct {
	OxdId string `json:"oxd_id"`
	Rpt string `json:"rpt"`
	Path string `json:"path"`
	HttpMethod string `json:"http_method"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

type RsCheckAccessResponseParams struct {
	
	Status string `json:"status"` 
	
	
	Access string `json:"access"`
	WwwAuthenticateHeader string `json:"www-authenticate_header"`
	Ticket string `json:"ticket"`
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`


}


