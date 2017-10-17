package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/token"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
	"github.com/juju/loggo"
	//"oxd-client-demo/utils"
)

func GetClientTokenPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars , globalvariables conf.GlobalVars) string {
	
	log :=loggo.GetLogger("default")
	var oxdResponse transport.OxdResponse
	var request model.GetClientTokenRequestParams
	request.OxdId = globalvariables.Oxdid
	request.ClientID = globalvariables.Clientid
	request.ClientSecret = globalvariables.Clientsecret
	request.OpHost = globalvariables.Ophost 
	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_CLIENT_TOKEN,request),
		&oxdResponse,
		globalvariables.Host)		 

	var response model.GetClientTokenResponseParams
	oxdResponse.GetParams(&response)
	var res = &response
	log.Debugf(res.AccessToken)
	return res.AccessToken
	
}
