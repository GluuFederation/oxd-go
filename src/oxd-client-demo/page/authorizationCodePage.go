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
	"oxd-client/constants"
	"oxd-client/model/params/token"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
	"oxd-client-demo/service"
)

func AuthorizationCodePageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	var oxdResponse transport.OxdResponse

	service.CallOxdServer(
		client.BuildOxdRequest(constants.GET_AUTHORIZATION_CODE,
			model.AuthorizationCodeRequestParams{session.OxdId,
				configuration.Username, configuration.Password, make([]string, 0),"",session.State}),
		&oxdResponse, transport.OxdConnectionParam{})

	var response model.AuthorizationCodeResponseParams
	oxdResponse.GetParams(&response)
	session.Code = response.Code
	utils.DisplayResponse(w,response)
}

