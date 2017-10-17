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

func RpGetClaimsGatheringUrlPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars,  accesstoken string,globalvariables conf.GlobalVars) {

   var oxdResponse transport.OxdResponse
   var requestParams uma.RpGetClaimsGatheringUrlRequestParams
   requestParams.OxdId = globalvariables.Oxdid
   requestParams.Ticket = session.UMATicket
   requestParams.ClaimsRedirectURI = "https://client.example.com"
   requestParams.ProtectionAccessToken = accesstoken
   
   page.CallOxdServer(
	client.BuildOxdRequest(constants.RP_GET_CLAIMS_GATHERING_URL,requestParams),
	&oxdResponse,
	globalvariables.Host)

	var response uma.RpGetClaimsGatheringUrlResponseParams
	oxdResponse.GetParams(&response)
	
	utils.DisplayResponse(w,response)




}