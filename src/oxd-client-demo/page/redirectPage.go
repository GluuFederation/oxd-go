//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/token"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
	"strings"
	
	
)
var serverConf conf.Configuration
var session conf.SessionVars
var State = " "
var Code = " "
var AccessToken = " "
var IdToken = " "
var RefreshToken = " "

func RedirectPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars,  accesstoken string,globalvariables conf.GlobalVars) conf.SessionVars  {
     
	readQueryParams(r,session)
	response := getTokensByCode(*session,configuration,accesstoken, globalvariables )
	saveTokens(response,session)
	getaccessrefreshtoken := getAccessTokenByRefreshToken(*session,configuration,accesstoken, globalvariables)
	saveRefreshTokens(getaccessrefreshtoken,session)
	getuserinfo := UserInfo(*session, configuration , accesstoken, globalvariables  )
	return getuserinfo
}

func readQueryParams(r *http.Request, session *conf.SessionVars){
	State = r.URL.Query().Get("state")
    Code = r.URL.Query().Get("code")
	
}

func getTokensByCode(session conf.SessionVars,configuration conf.Configuration,  accesstoken string,globalvariables conf.GlobalVars) model.TokensByCodeResponseParams{
	
	var oxdResponse transport.OxdResponse
	ConnectionType := globalvariables.ConnectionType
	HttpRestUrl := globalvariables.Httpresturl 
	if(ConnectionType == "local") {
	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_TOKENS_BY_CODE,
			model.TokensByCodeRequestParams{OxdId : globalvariables.Oxdid, ProtectionAccessToken : accesstoken, Code: Code, State : State }),
		&oxdResponse,
		globalvariables.Host)
		} else {
			page.CallOxdHttpsExtension(
				client.BuildOxdRequest(constants.GET_TOKENS_BY_CODE,
					model.TokensByCodeRequestParams{OxdId : globalvariables.Oxdid, ProtectionAccessToken : accesstoken, Code: Code, State : State }),
				&oxdResponse,
				HttpRestUrl)
		}
	var response model.TokensByCodeResponseParams
	oxdResponse.GetParams(&response)
	return response
}

func saveTokens(response model.TokensByCodeResponseParams, session *conf.SessionVars){

	
	AccessToken = response.AccessToken
	IdToken = response.IdToken
	RefreshToken = response.RefreshToken
	
	
}

func saveRefreshTokens(response model.GetAccessTokenByRefreshTokenResponseParams, session *conf.SessionVars){
	AccessToken = response.AccessToken
	RefreshToken= response.RefreshToken
	
}

func getAccessTokenByRefreshToken(session conf.SessionVars,configuration conf.Configuration,  accesstoken string,globalvariables conf.GlobalVars) model.GetAccessTokenByRefreshTokenResponseParams{
	var oxdResponse transport.OxdResponse
	ConnectionType := globalvariables.ConnectionType
	HttpRestUrl := globalvariables.Httpresturl
	if(ConnectionType == "local") {
		page.CallOxdServer(
			client.BuildOxdRequest(constants.GET_ACCESS_TOKEN_BY_REFRESH_TOKEN,
				model.GetAccessTokenByRefreshTokenRequestParams{OxdId:globalvariables.Oxdid , RefreshToken: RefreshToken , ProtectionAccessToken:accesstoken}),

			&oxdResponse,
			globalvariables.Host)
			} else {
				page.CallOxdHttpsExtension(
					client.BuildOxdRequest(constants.GET_ACCESS_TOKEN_BY_REFRESH_TOKEN,
						model.GetAccessTokenByRefreshTokenRequestParams{OxdId:globalvariables.Oxdid , RefreshToken: RefreshToken , ProtectionAccessToken:accesstoken}),
		
					&oxdResponse,
					HttpRestUrl)
                   
			}
	
		var response model.GetAccessTokenByRefreshTokenResponseParams
		oxdResponse.GetParams(&response)
		return response
	}


	func UserInfo( session conf.SessionVars, configuration conf.Configuration, accesstoken string,globalvariables conf.GlobalVars)  conf.SessionVars   {

			var oxdResponse transport.OxdResponse
		    ConnectionType := globalvariables.ConnectionType
			HttpRestUrl := globalvariables.Httpresturl
			if(ConnectionType == "local") {
			page.CallOxdServer(
				client.BuildOxdRequest(constants.GET_USER_INFO,model.UserInfoRequestParams{OxdId : globalvariables.Oxdid, ProtectionAccessToken : accesstoken, AccessToken : AccessToken}),
				&oxdResponse,
				globalvariables.Host)
			} else {
				page.CallOxdHttpsExtension(
					client.BuildOxdRequest(constants.GET_USER_INFO,model.UserInfoRequestParams{OxdId : globalvariables.Oxdid, ProtectionAccessToken : accesstoken, AccessToken : AccessToken}),
					&oxdResponse,
					HttpRestUrl)
			}
			var response model.UserInfoResponseParams
			oxdResponse.GetParams(&response)
			username := response.Claims["name"]
			email := response.Claims["email"]
			usernameString := strings.Join(username," ")
			emailString := strings.Join(email," ")
			session.Name = usernameString
			session.Email = emailString
			return session
		

		}