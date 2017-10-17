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
	"oxd-client/model/params/registration"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
)

func RegisterSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	var oxdResponse transport.OxdResponse


	page.CallOxdServer(
		client.BuildOxdRequest(constants.REGISTER_SITE,configuration.RegisterSiteRequestParams),
		&oxdResponse,
		configuration.Host)

	var response model.RegisterSiteResponseParams
	oxdResponse.GetParams(&response)
	session.OxdId = response.OxdId
	utils.DisplayResponse(w,response)
}



