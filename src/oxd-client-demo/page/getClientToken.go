package page

import (
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/params/token"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
)

func GetClientTokenPageSite(configuration *conf.Configuration) string {

	requestParams := model.GetClientTokenRequestParams{configuration.RegisterSiteParams.OxdId,
		configuration.ClientId,
		configuration.ClientSecret,
		configuration.IdpHost,
		"", []string{"uma_protection"}}

	requestType := transport.GetRequestTypeByName(configuration.Type)
	address := utils.GetAddressForType(*configuration, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_CLIENT_TOKEN, requestParams)

	var response transport.OxdResponse
	var responseParams model.GetClientTokenResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	configuration.AccessToken = responseParams.AccessToken
	return responseParams.AccessToken

}
