//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"oxd-client-demo/conf"
	// "oxd-client-demo/service"
	// "oxd-client/client"
	// "oxd-client/constants"
	"oxd-client/model/params"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
)

func UserInfoPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars, userinfo string) {

	var oxdResponse transport.OxdResponse

	// page.CallOxdServer(
	// 	client.BuildOxdRequest(constants.GET_USER_INFO,model.UserInfoRequestParams{session.OxdId,session.AccessToken}),
	// 	&oxdResponse,
	// 	configuration.Host)

	var response model.UserInfoResponseParams
	oxdResponse.GetParams(&response)
	utils.DisplayResponse(w,response)
}



