package page

import (
	"fmt"
	"net/http"
    "oxd-client-demo/conf"
	"oxd-client/model/params/registration"
	"oxd-client/client"
    "oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/oxd_config_json"
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
)

  
func SetupClientPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars, globalvariables conf.GlobalVars)   {
	
	 
	var oxdResponse transport.OxdResponse
	var request model.SetupClientRequestParams
	
	request.Scope = globalvariables.Scope
	request.OpHost = r.FormValue("OphostId")
	request.ClientId = r.FormValue("Clientid")
	request.GrantType = globalvariables.GrantType
	request.ClientName = r.FormValue("Client_name")
	request.ClientSecret = r.FormValue("Clientsecret")
	request.AuthorizationRedirectUri = r.FormValue("RedirectUrl")
	request.PostLogoutRedirectUri = r.FormValue("PostLogoutRedirectUrl")
	OxdHost := fmt.Sprint(configuration.Host+":"+r.FormValue("OxdPort"))
	ConnectionType := r.FormValue("Connectiontype")
	HttpRestUrl := r.FormValue("Httpresturl")
	

	if(ConnectionType == "local") {
			page.CallOxdServer(
				client.BuildOxdRequest(constants.SETUP_CLIENT,request),
				&oxdResponse,
				OxdHost)
	} else {
			page.CallOxdHttpsExtension(
				client.BuildOxdRequest(constants.SETUP_CLIENT,request),
				&oxdResponse,
				HttpRestUrl)

			}

	var response model.SetupClientResponseParams
	
	var savesettings newmodel.OxdSetting
     oxdResponse.GetParams(&response)
	 
	session.OxdId = response.OxdId
	utils.DisplayResponse(w,response)
	
	savesettings.OxdId = response.OxdId
	savesettings.ClientId = response.ClientId
	savesettings.OxdHost = configuration.Host
	savesettings.ClientName = request.ClientName
	savesettings.OxdPort = r.FormValue("OxdPort")
	savesettings.ClientSecret = response.ClientSecret
	savesettings.ConnectionType = r.FormValue("Connectiontype")
	savesettings.PostLogoutRedirectUri = request.PostLogoutRedirectUri
	savesettings.Scope = configuration.RegisterSiteRequestParams.Scope
	savesettings.OpHost = request.OpHost 
	savesettings.AuthorizationRedirectUri = request.AuthorizationRedirectUri
	savesettings.GrantType = configuration.RegisterSiteRequestParams.GrantType
	savesettings.ResponseTypes = configuration.RegisterSiteRequestParams.ResponseTypes
    savesettings.HttpRestUrl = HttpRestUrl
	utils.SaveoxdSetting(w,savesettings)
	
    
	
	
}