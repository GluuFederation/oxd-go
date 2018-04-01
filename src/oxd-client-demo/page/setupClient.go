package page

import (
	//"fmt"
	"net/http"
    "oxd-client-demo/conf"
	"oxd-client/model/params/registration"
	"oxd-client/client"
    "oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
)

  
func SetupClientPage(w http.ResponseWriter, r *http.Request, conf *conf.Configuration) string  {
	requestParams := prepareSetupParams(*conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.SETUP_CLIENT, requestParams)

	var response transport.OxdResponse
	var responseParams model.SetupClientResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.ClientOxdId = responseParams.OxdId
	return responseParams.OxdId
}

func prepareSetupParams(conf conf.Configuration)  model.SetupClientRequestParams {
	var requestParams model.SetupClientRequestParams
	requestParams.OpHost = conf.IdpHost
	requestParams.AuthorizationRedirectUri = conf.RegisterSiteParams.RedirectUris[0]
	requestParams.PostLogoutRedirectUri = conf.RegisterSiteParams.PostLogoutRedirectUri
	requestParams.ClientLogoutUri = conf.RegisterSiteParams.ClientLogoutUri
	requestParams.RedirectUris = conf.RegisterSiteParams.RedirectUris
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = conf.RegisterSiteParams.Scope
	requestParams.GrantTypes = conf.RegisterSiteParams.GrantTypes
	requestParams.ResponseTypes = conf.RegisterSiteParams.ResponseTypes
	return  requestParams
}