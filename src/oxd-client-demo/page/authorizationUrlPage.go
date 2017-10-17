//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	//"fmt"
	
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/url"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
	"github.com/juju/loggo"
)

func AuthorizationUrlPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars,  accesstoken string ,  globalvariables conf.GlobalVars) {

	log :=loggo.GetLogger("default")
	var oxdResponse transport.OxdResponse
	row := model.AuthorizationUrlRequestParams{
        CustomParameters: make(map[string]string),
    }

	row.CustomParameters["param1"] = "value1"
	row.CustomParameters["param2"] = "value2"
	row.OxdId = globalvariables.Oxdid
	row.ProtectionAccessToken = accesstoken

	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_AUTHORIZATION_URL,
			row),
		&oxdResponse,
		globalvariables.Host)


	var response model.AuthorizationUrlResponseParams
	oxdResponse.GetParams(&response)
	 
	log.Debugf(response.AuthorizationUrl)
	
	http.Redirect(w, r, response.AuthorizationUrl, http.StatusPermanentRedirect)
	
}
