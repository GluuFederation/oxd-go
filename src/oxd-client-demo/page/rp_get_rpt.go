package page

import (
	
	
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/uma"
	"oxd-client/constants"
	"oxd-client/model/transport"
	
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
)

func RpGetRptPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars, accesstoken string,globalvariables conf.GlobalVars) {

   var oxdResponse transport.OxdResponse
   var requestParams uma.RpGetRptRequestParams

    requestParams.OxdId = globalvariables.Oxdid
    requestParams.Ticket = session.UMATicket 
	requestParams.ProtectionAccessToken = accesstoken
	ConnectionType := globalvariables.ConnectionType
	HttpRestUrl := globalvariables.Httpresturl 
	if(ConnectionType == "local") {
    page.CallOxdServer(
	client.BuildOxdRequest(constants.RP_GET_RPT,requestParams),
	&oxdResponse,
	globalvariables.Host)
	} else {
		page.CallOxdHttpsExtension(
			client.BuildOxdRequest(constants.RP_GET_RPT,requestParams),
			&oxdResponse,
			HttpRestUrl)
	}
	var response uma.RpGetRptResponseParams
	oxdResponse.GetParams(&response)
	
	utils.DisplayResponse(w,response)




}