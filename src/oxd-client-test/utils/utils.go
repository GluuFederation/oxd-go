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
	urlParam "oxd-client/model/params/url"
	token "oxd-client/model/params/token"
	"oxd-client/model/params/registration"
    //"github.com/pkg/term"
	//"strings"
	//"gopkg.in/resty.v1"
	//"encoding/json"
	//"strings"
	//"github.com/go-resty/resty"
	//"encoding/json"
	//"strings"
	//"github.com/go-resty/resty"
	//"encoding/json"
	//"strings"
	//"bufio"
	//"os"
	"fmt"
	"net/url"
	"strings"
	"github.com/go-resty/resty"
	"crypto/tls"
	//"net/http"
	//"net/http"
	"net/http"
	"net/http/cookiejar"
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
	requestParams.ClientRegistrationClientUri = "https://idp.example.com/identity/"
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

func ExecCodeFlow() (urlParam.AuthorizationCodeFlowResponseParams,string){
	requestParams := urlParam.AuthorizationCodeFlowRequestParams{
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
	var responseParams urlParam.AuthorizationCodeFlowResponseParams
	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams, requestParams.OxdId
}


func GetCode(getRequest GetRequest) (url.Values) {
	authorizationUrl, authParam := GetAuthorizationUrl(GetRestRequest)
	urlParam := strings.Replace(authorizationUrl,"restv1/","",-1)
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify:true})
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(100))
	resty.SetDebug(true)
	resty.SetCookie(&http.Cookie{
		Name:"csfcfc",
		Value:"5O2CVUpfqL73",
		Path: "/",
		Domain: "idp.example.com",
		MaxAge: 36000,
		HttpOnly: true,
		Secure: false,
	})
	jar,_ := cookiejar.New(nil)
	resty.SetCookieJar(jar)
	resp,_ := resty.R().Get(urlParam)
	fmt.Print(resp.Request.URL)


	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify:true})
	resty.SetRedirectPolicy(resty.NoRedirectPolicy())
    resty.SetCookieJar(jar)
    resp2,_ :=resty.R().SetFormData(map[string]string{
		"loginForm": "loginForm",
		"loginForm:username": "admin",
		"loginForm:password": "test123",
		"loginForm:loginButton": "Login",
		"javax.faces.ViewState": "stateless",
	}).Post("https://idp.example.com/oxauth/login")

	resp3,_ := resty.R().Get(resp2.RawResponse.Header.Get("Location"))
	resp4,_ := resty.R().Get(resp3.RawResponse.Header.Get("Location"))
	redirectUrl,_ := url.Parse(resp4.RawResponse.Header.Get("Location"))
	values := redirectUrl.Query()
	values.Add("oxdId",authParam.OxdId)

	return values
}

func PrepareAuthorizationUrlRequestParams(getRequest GetRequest) urlParam.AuthorizationUrlRequestParams {
	return urlParam.AuthorizationUrlRequestParams{RegisterClientSite(getRequest),
		"",
		make([]string,0),
		[]string{"openid"},
		"",
		nil}
}

func GetAuthorizationUrl(getRequest GetRequest) (string, urlParam.AuthorizationUrlRequestParams) {
	//BEFORE
	requestParams := PrepareAuthorizationUrlRequestParams(getRequest)
	request, connectionParam := getRequest(constants.GET_AUTHORIZATION_URL, requestParams)

	var response transport.OxdResponse
	var responseParams urlParam.AuthorizationUrlResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	return responseParams.AuthorizationUrl, requestParams
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
	requestParams := PrepareRegisterParams(GetClientToken(getRequest))
	request, connectionParam := getRequest(constants.REGISTER_SITE,requestParams)

	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams.OxdId
}