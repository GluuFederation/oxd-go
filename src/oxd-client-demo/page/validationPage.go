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
	"oxd-client/model/params/validation"
	"oxd-client-demo/utils"
)

func ValidationPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars) {
	validateAccessToken(w,configuration,session)
	validateIdToken(w,configuration, session)
}

func validateIdToken(w http.ResponseWriter, configuration conf.Configuration, session conf.SessionVars){
	var oxdResponse transport.OxdResponse

	service.CallOxdServer(
		client.BuildOxdRequest(constants.CHECK_ID_TOKEN,
			validation.CheckIdTokenRequestParams{session.OxdId,
				session.IdToken}),
		&oxdResponse,
		transport.OxdConnectionParam{})

	var response validation.CheckIdTokenResponseParams
	oxdResponse.GetParams(&response)
	utils.DisplayResponse(w,response)
}

func validateAccessToken(w http.ResponseWriter, configuration conf.Configuration, session conf.SessionVars){
	var oxdResponse transport.OxdResponse

	service.CallOxdServer(
		client.BuildOxdRequest(constants.CHECK_ACCESS_TOKEN,
			validation.CheckAccessTokenRequestParams{session.OxdId,
				session.IdToken, session.AccessToken}),
		&oxdResponse,
		transport.OxdConnectionParam{})

	var response validation.CheckAccessTokenResponseParams
	oxdResponse.GetParams(&response)
	utils.DisplayResponse(w,response)
}
