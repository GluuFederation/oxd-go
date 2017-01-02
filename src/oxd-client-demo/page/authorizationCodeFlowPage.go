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
	"oxd-client/model/params/url"
	"oxd-client/constants"
	"oxd-client-demo/service"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
)

func AuthorizationCodeFlowPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars) {
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.AUTHORIZATION_CODE_FLOW,
			model.AuthorizationCodeFlowRequestParams{session.OxdId,"https://10.0.0.250:8080/redirect","@!BB85.C095.8C1E.0573!0001!F554.249C!0008!5322.3303","","","","openid","",""}),
		&oxdResponse,
		configuration.Host)
	var response model.AuthorizationCodeFlowResponseParams
	oxdResponse.GetParams(&response)
	utils.DisplayResponse(w,response)
}


