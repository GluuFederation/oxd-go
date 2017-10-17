package page

import (
	"fmt"
	"net/http"
    "oxd-client-demo/conf"
	"oxd-client/model/params/registration"
	"oxd-client/client"
    "oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client/model/oxd_config_json"
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
)

  
func SetupClientPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars )   {
	
	 
	var oxdResponse transport.OxdResponse
	var request model.SetupClientRequestParams

	request.OpHost = r.FormValue("OphostId")
	request.ClientId = r.FormValue("clientid")
	request.ClientName = r.FormValue("Client_name")
	request.ClientSecret = r.FormValue("clientsecret")
	request.AuthorizationRedirectUri = r.FormValue("RedirectUrl")
	request.Scope = configuration.RegisterSiteRequestParams.Scope
	request.PostLogoutRedirectUri = r.FormValue("PostLogoutRedirectUrl")
	OxdHost := fmt.Sprint(configuration.Host+":"+r.FormValue("OxdPort"))
	request.GrantType = configuration.RegisterSiteRequestParams.GrantType
	
	page.CallOxdServer(
		client.BuildOxdRequest(constants.SETUP_CLIENT,request),
		&oxdResponse,
		OxdHost)

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
	savesettings.OpHost = configuration.RegisterSiteRequestParams.OpHost
	savesettings.AuthorizationRedirectUri = request.AuthorizationRedirectUri
	savesettings.GrantType = configuration.RegisterSiteRequestParams.GrantType
	savesettings.ResponseTypes = configuration.RegisterSiteRequestParams.ResponseTypes

	utils.SaveoxdSetting(w,savesettings)
	
    
	
	
}