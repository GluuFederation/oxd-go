//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package service

import (
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/url"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
	"net/http"
	token "oxd-client/model/params/token"
)

func GetAuthorizationUrlPage(conf *conf.Configuration) string {

	requestParams := PrepareAuthorizationUrlRequestParams(conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_AUTHORIZATION_URL, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.AuthorizationUrlResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.OpenIdConnect.AuthorizationUrl = responseParams.AuthorizationUrl
	return responseParams.AuthorizationUrl
}

func PrepareAuthorizationUrlRequestParams(conf *conf.Configuration) model.AuthorizationUrlRequestParams {
	return model.AuthorizationUrlRequestParams{conf.RegisterSiteParams.OxdId,
		"",
		make([]string,0),
		 conf.RegisterSiteParams.Scope,
		"",
		nil}
}

func ReadCode(w http.ResponseWriter, r *http.Request, conf *conf.Configuration)  {
	code := r.URL.Query()["code"]
	state := r.URL.Query()["state"]
	conf.OpenIdConnect.Code = code[0]
	conf.OpenIdConnect.State = state[0]
	http.Redirect(w,r,"/",http.StatusMovedPermanently)
}

func GetTokenWithCode( conf *conf.Configuration) string {
	requestParams := PrepareTokenRequestParams(conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_TOKENS_BY_CODE, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams token.TokensByCodeResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.OpenIdConnect.IdToken = responseParams.IdToken
	conf.OpenIdConnect.RefreshToken = responseParams.RefreshToken
	conf.OpenIdConnect.AccessToken = responseParams.AccessToken

	return responseParams.IdToken
}

func GetTokenWithRefreshToken( conf *conf.Configuration) string {
	requestParams := PrepareRefreshRequestParams(conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_ACCESS_TOKEN_BY_REFRESH_TOKEN, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams token.GetAccessTokenByRefreshTokenResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.OpenIdConnect.AccessToken = responseParams.AccessToken
	conf.OpenIdConnect.RefreshToken = responseParams.RefreshToken

	return responseParams.AccessToken
}
func PrepareRefreshRequestParams(conf *conf.Configuration) token.GetAccessTokenByRefreshTokenRequestParams {
	return token.GetAccessTokenByRefreshTokenRequestParams{conf.RegisterSiteParams.OxdId,
		conf.OpenIdConnect.RefreshToken,
		"",
		conf.RegisterSiteParams.Scope}
}

func PrepareTokenRequestParams(conf *conf.Configuration) token.TokensByCodeRequestParams {
	return token.TokensByCodeRequestParams{conf.RegisterSiteParams.OxdId,
		"",
		conf.OpenIdConnect.Code,
		conf.OpenIdConnect.State}
}

func GetUserInfo( conf *conf.Configuration) map[string] []string {
	requestParams := prepareUserInfoRequestParams(conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_USER_INFO, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams token.UserInfoResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.OpenIdConnect.Claims = responseParams.Claims

	return responseParams.Claims
}
func prepareUserInfoRequestParams(conf *conf.Configuration) token.UserInfoRequestParams {
	return token.UserInfoRequestParams{
		conf.RegisterSiteParams.OxdId,
		"",
		conf.OpenIdConnect.AccessToken}
}

func GetLogoutUrl( conf *conf.Configuration) string {
	requestParams := prepareLogoutUrlRequestParams(conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.GET_LOGOUT_URI, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.LogoutUrlResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.OpenIdConnect.LogoutUrl = responseParams.Uri

	return responseParams.Uri
}
func prepareLogoutUrlRequestParams(conf *conf.Configuration) model.LogoutUrlRequestParams {
	return model.LogoutUrlRequestParams{conf.RegisterSiteParams.OxdId,
		"",
		conf.OpenIdConnect.IdToken,
		conf.RegisterSiteParams.PostLogoutRedirectUri,
		"",
		""}
}