package page

import (
	"oxd-client/constants"
	"oxd-client/client"
	"oxd-client/model/params/registration"
	"oxd-client-demo/conf"
	"oxd-client-demo/utils"
	"oxd-client/model/transport"
	"net/http"
)


func RegisterClientSite(conf *conf.Configuration) string{
	requestParams := prepareRegisterParams(*conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.REGISTER_SITE, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.RegisterSiteParams.OxdId = responseParams.OxdId

	GetClientTokenPageSite(conf)

	return responseParams.OxdId;
}

func prepareRegisterParams(conf conf.Configuration) model.RegisterSiteRequestParams{
	var requestParams model.RegisterSiteRequestParams
	requestParams.OpHost = conf.IdpHost
	requestParams.AuthorizationRedirectUri = conf.RegisterSiteParams.AuthorizationRedirectUri
	requestParams.PostLogoutRedirectUri = conf.RegisterSiteParams.PostLogoutRedirectUri
	requestParams.ClientLogoutUri = conf.RegisterSiteParams.ClientLogoutUri
	requestParams.RedirectUris = conf.RegisterSiteParams.RedirectUris
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = conf.RegisterSiteParams.Scope
	requestParams.GrantTypes = conf.RegisterSiteParams.GrantTypes
	requestParams.ResponseTypes = conf.RegisterSiteParams.ResponseTypes
	requestParams.ClientId = conf.ClientId
	requestParams.ClientSecret = conf.ClientSecret
	requestParams.ClientRegistrationClientUri = conf.ClientRegistrationClientUri
	return  requestParams
}

func UpdateClientSite(w http.ResponseWriter, r *http.Request, conf *conf.Configuration) string{

	utils.ParseResponse(w,r,conf.RegisterSiteParams)

	requestParams := prepareUpdateParams(*conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.UPDATE_SITE, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.UpdateSiteResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	return responseParams.OxdId
}

func prepareUpdateParams(conf conf.Configuration)  model.UpdateSiteRequestParams {
	var requestParams model.UpdateSiteRequestParams
	requestParams.OxdId = conf.RegisterSiteParams.OxdId
	requestParams.AuthorizationRedirectUri = conf.RegisterSiteParams.AuthorizationRedirectUri
	requestParams.PostLogoutRedirectUri = conf.RegisterSiteParams.PostLogoutRedirectUri
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = conf.RegisterSiteParams.Scope
	requestParams.GrantTypes = conf.RegisterSiteParams.GrantTypes
	requestParams.ResponseTypes = conf.RegisterSiteParams.ResponseTypes
	requestParams.ClientId = conf.ClientId
	requestParams.ClientSecret = conf.ClientSecret
	return  requestParams
}

func RemoveClientSite(conf *conf.Configuration) string{
	requestParams := model.RemoveSiteRequestParams{conf.RegisterSiteParams.OxdId}

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.REMOVE_SITE, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.RemoveSiteResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.RegisterSiteParams.OxdId = ""
	return ""
}

