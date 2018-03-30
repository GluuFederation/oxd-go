//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client-demo/service"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client/model/params/url"
)

func LogoutUrlPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars, accesstoken string,globalvariables conf.GlobalVars) {
	var oxdResponse transport.OxdResponse
	ConnectionType := globalvariables.ConnectionType
	//HttpRestUrl := globalvariables.Httpresturl
	if(ConnectionType == "local") {
		service.CallOxdServer(
		client.BuildOxdRequest(constants.GET_LOGOUT_URI,  
			model.LogoutUrlRequestParams{OxdId: session.OxdId,ProtectionAccessToken : accesstoken, IdTokenHint: session.IdToken, PostLogoutRedirectUri: "",State: session.State,SessionState: ""}),
		&oxdResponse,
		transport.OxdConnectionParam{})
	} else {
        //page.CallOxdHttpsExtension(
			//client.BuildOxdRequest(constants.GET_LOGOUT_URI,
			//	model.LogoutUrlRequestParams{OxdId: session.OxdId,ProtectionAccessToken : accesstoken, IdTokenHint: session.IdToken, PostLogoutRedirectUri: "",State: session.State,SessionState: ""}),
			//&oxdResponse,
			//HttpRestUrl)
	}
	var response model.LogoutUrlResponseParams
	oxdResponse.GetParams(&response)
	http.Redirect(w, r, response.Uri, http.StatusPermanentRedirect)
}
