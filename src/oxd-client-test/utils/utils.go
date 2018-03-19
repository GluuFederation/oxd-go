//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package utils

import (
	"oxd-client/client"
	"oxd-client/constants"

	"oxd-client-test/conf"
	"oxd-client/model/transport"
	url "oxd-client/model/params/url"
	token "oxd-client/model/params/token"
	"oxd-client/model/params/registration"
	//"testing"
	//"github.com/stretchr/testify/assert"
	//"testing"
	//"github.com/stretchr/testify/assert"
	//"testing"
	//"github.com/stretchr/testify/assert"
)

type GetRequest func(command constants.CommandType, params transport.Param) (transport.OxdRequest, transport.OxdConnectionParam)


func RegisterClient(accessToken string) string{
	requestParams := PrepareRegisterParams(accessToken)
	request := client.BuildOxdRequest(constants.REGISTER_SITE,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.REGISTER_SITE}
	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams.OxdId
}

func SetupClient() model.SetupClientResponseParams{
	requestParams := PrepareSetupParams()

	request, connectionParam := GetRestRequest(constants.SETUP_CLIENT,requestParams)
	var response transport.OxdResponse
	var responseParams model.SetupClientResponseParams

	client.Send(request,connectionParam,&response)
	response.GetParams(&responseParams)

	conf.TestConfiguration.OpHost = responseParams.OpHost
	conf.TestConfiguration.ClientId = responseParams.ClientId
	conf.TestConfiguration.ClientSecret = responseParams.ClientSecret
	conf.TestConfiguration.OxdId = responseParams.OxdId

	return responseParams
}

func PrepareRegisterParams(accessToken string)  model.RegisterSiteRequestParams {
	var requestParams model.RegisterSiteRequestParams
	requestParams.OpHost = conf.TestConfiguration.OpHost
	requestParams.AuthorizationRedirectUri = conf.TestConfiguration.RedirectUrl
	requestParams.PostLogoutRedirectUri = conf.TestConfiguration.PostLogoutRedirectUrl
	requestParams.ClientLogoutUri = []string {conf.TestConfiguration.LogoutUrl}
	requestParams.RedirectUris = [] string {conf.TestConfiguration.RedirectUrl}
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = []string{"openid", "profile","uma_protection","uma_authorization"}
	requestParams.GrantTypes = []string{"authorization_code"}
	requestParams.ResponseTypes = []string{"code"}
	requestParams.ClientId = conf.TestConfiguration.ClientId
	requestParams.ClientSecret = conf.TestConfiguration.ClientSecret
	requestParams.ClientRegistrationClientUri = "https://localhost/identity/"
	//requestParams.ClientRegistrationAccessToken = accessToken
	//requestParams.ProtectionAccessToken = accessToken
	return  requestParams
}

func PrepareSetupParams()  model.SetupClientRequestParams {
	var requestParams model.SetupClientRequestParams
	requestParams.OpHost = conf.TestConfiguration.OpHost
	requestParams.AuthorizationRedirectUri = conf.TestConfiguration.RedirectUrl
	requestParams.PostLogoutRedirectUri = conf.TestConfiguration.PostLogoutRedirectUrl
	requestParams.ClientLogoutUri = []string {conf.TestConfiguration.LogoutUrl}
	requestParams.RedirectUris = [] string {conf.TestConfiguration.RedirectUrl}
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = []string{"openid", "profile","uma_protection","uma_authorization"}
	requestParams.GrantTypes = []string{"authorization_code"}
	requestParams.ResponseTypes = []string{"code"}
	return  requestParams
}

func PrepareUpdateParams()  model.UpdateSiteRequestParams {
	setupResponse := SetupClient()
	var requestParams model.UpdateSiteRequestParams
	requestParams.OxdId = setupResponse.OxdId
	requestParams.AuthorizationRedirectUri = conf.TestConfiguration.RedirectUrl
	requestParams.PostLogoutRedirectUri = conf.TestConfiguration.PostLogoutRedirectUrl
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = []string{"openid", "profile","uma_protection","uma_authorization"}
	requestParams.GrantType = []string{"authorization_code"}
	requestParams.ResponseTypes = []string{"code"}
	requestParams.ClientId = setupResponse.ClientId
	requestParams.ClientSecret = setupResponse.ClientSecret
	//requestParams.ProtectionAccessToken = GetClientToken(GetRestRequest)
	return  requestParams
}

func ExecCodeFlow() (url.AuthorizationCodeFlowResponseParams,string){
	requestParams := url.AuthorizationCodeFlowRequestParams{
		RegisterClient(""),
		conf.TestConfiguration.RedirectUrl,
		conf.TestConfiguration.ClientId,
		conf.TestConfiguration.ClientSecret,
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		conf.TestConfiguration.Scope,"",""}
	request := client.BuildOxdRequest(constants.AUTHORIZATION_CODE_FLOW,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.AUTHORIZATION_CODE_FLOW}

	var response transport.OxdResponse
	var responseParams url.AuthorizationCodeFlowResponseParams
	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams, requestParams.OxdId
}


func ExecCode() (token.AuthorizationCodeResponseParams, token.AuthorizationCodeRequestParams) {
	requestParams := token.AuthorizationCodeRequestParams{RegisterClient(""),
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		make([]string,0),
		"","af0ifjsldkj"}
	request := client.BuildOxdRequest(constants.GET_AUTHORIZATION_CODE,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.GET_AUTHORIZATION_CODE}

	var response transport.OxdResponse
	var responseParams token.AuthorizationCodeResponseParams
	client.Send(request,connectionParam,&response)
	response.GetParams(&responseParams)
	return responseParams, requestParams
}

func GetSocketRequest (command constants.CommandType, params transport.Param) (transport.OxdRequest, transport.OxdConnectionParam){
	request := client.BuildOxdRequest(command,params)
	connectionParam := transport.OxdConnectionParam {conf.TestConfiguration.Host,
	transport.SOCKET,
	"",
	command}
	return request, connectionParam
}

func GetRestRequest (command constants.CommandType, params transport.Param) (transport.OxdRequest, transport.OxdConnectionParam){
	var accesstoken = ""
	if(command != constants.GET_CLIENT_TOKEN && command != constants.SETUP_CLIENT){
		accesstoken = GetClientToken(GetRestRequest)
	}
	request := client.BuildOxdRequest(command,params)
	connectionParam := transport.OxdConnectionParam {conf.TestConfiguration.OxdHttpsHost,
	transport.REST,
	accesstoken,
	command}

	return request, connectionParam
}

func GetClientToken(getRequest GetRequest) string{

	requestParams := token.GetClientTokenRequestParams {conf.TestConfiguration.OxdId,
		conf.TestConfiguration.ClientId,
		conf.TestConfiguration.ClientSecret,
		conf.TestConfiguration.OpHost,
		"",[]string{"uma_protection"}}

	request, connectionParam := getRequest(constants.GET_CLIENT_TOKEN,requestParams)
	var response transport.OxdResponse
	var responseParams token.GetClientTokenResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams.AccessToken
}


func RegisterClientSite(getRequest GetRequest) string{
	//BEFORE
	requestParams := PrepareRegisterParams(GetClientToken(getRequest))
	request, connectionParam := getRequest(constants.REGISTER_SITE,requestParams)

	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	return responseParams.OxdId
}