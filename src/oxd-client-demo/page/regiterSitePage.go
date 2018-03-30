//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"oxd-client-demo/conf"
)

func RegisterSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	//var oxdResponse transport.OxdResponse

	//
	//service.CallOxdServer(
	//	client.BuildOxdRequest(constants.REGISTER_SITE,configuration.RegisterSiteRequestParams),
	//	&oxdResponse,
	//	transport.OxdConnectionParam{})
	//
	//var response model.RegisterSiteResponseParams
	//oxdResponse.GetParams(&response)
	//session.OxdId = response.OxdId
	//utils.DisplayResponse(w,response)
}



