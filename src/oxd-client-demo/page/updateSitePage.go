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
	"oxd-client-demo/oxd_config_json"
	"oxd-client/model/params/registration"
	//"oxd-client-demo/utils"
)


func UpdateSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars, accesstoken string, globalvariables conf.GlobalVars) {
	var oxdResponse transport.OxdResponse
	
	var request model.UpdateSiteRequestParams

	request.OxdId = r.FormValue("OxdID")
	request.OpHost = r.FormValue("OphostId")
	request.ClientId = r.FormValue("Clientid")
	request.ProtectionAccessToken  = accesstoken
	request.ClientName = r.FormValue("Client_name")
	request.ClientSecret = r.FormValue("Clientsecret")
	request.AuthorizationRedirectUri = r.FormValue("RedirectUrl")
	request.Scope = globalvariables.Scope
	request.PostLogoutRedirectUri = r.FormValue("PostLogoutRedirectUrl")
	request.GrantType = globalvariables.GrantType
	ConnectionType := r.FormValue("Connectiontype")
	HttpRestUrl := r.FormValue("Httpresturl")
	
	if(ConnectionType == "local") {
    page.CallOxdServer(
		client.BuildOxdRequest(constants.UPDATE_SITE,request),
		&oxdResponse,
		globalvariables.Host)
	} else{
		page.CallOxdHttpsExtension(
			client.BuildOxdRequest(constants.UPDATE_SITE,request),
			&oxdResponse,
			HttpRestUrl)
	}

    var response model.UpdateSiteResponseParams
	var updatesettings newmodel.OxdSetting

	updatesettings.OxdHost = configuration.Host
	updatesettings.OxdId = request.OxdId
	updatesettings.PostLogoutRedirectUri = request.PostLogoutRedirectUri
	updatesettings.ResponseTypes = configuration.UpdateSiteRequestParams.ResponseTypes
	updatesettings.ClientName = request.ClientName
	updatesettings.Scope = configuration.UpdateSiteRequestParams.Scope
	updatesettings.OpHost = request.OpHost
	updatesettings.GrantType = configuration.UpdateSiteRequestParams.GrantType
	updatesettings.AuthorizationRedirectUri = request.AuthorizationRedirectUri
	updatesettings.ClientId = request.ClientId
	updatesettings.ClientSecret = request.ClientSecret
	oxdResponse.GetParams(&response)
	session.OxdId = response.OxdId
	//utils.UpdateoxdSetting(w,updatesettings)
	
}
