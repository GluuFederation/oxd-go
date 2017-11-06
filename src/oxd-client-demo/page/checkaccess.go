package page

import (
	
	
	"net/http"
	"fmt"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/uma"
	"oxd-client/constants"
	"oxd-client/model/transport"
	
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
)

func CheckAccessPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars,  accesstoken string, globalvariables conf.GlobalVars) {

   var oxdResponse transport.OxdResponse
   var requestParams uma.RsCheckAccessRequestParams
   requestParams.OxdId = globalvariables.Oxdid
   requestParams.Path = "/scim" 
   requestParams.Rpt = " "
   requestParams.HttpMethod = "GET"
   requestParams.ProtectionAccessToken = accesstoken
   ConnectionType := globalvariables.ConnectionType
   HttpRestUrl := globalvariables.Httpresturl 
   if(ConnectionType == "local") {
   page.CallOxdServer(
	client.BuildOxdRequest(constants.RS_CHECK_ACCESS,requestParams),
	&oxdResponse,
	globalvariables.Host)
   } else {
	page.CallOxdHttpsExtension(
		client.BuildOxdRequest(constants.RS_CHECK_ACCESS,
			requestParams),
		&oxdResponse,
		HttpRestUrl)
   }
	var response uma.RsCheckAccessResponseParams
	oxdResponse.GetParams(&response)
	session.UMATicket= response.Ticket
    fmt.Println(session.UMATicket)
	utils.DisplayResponse(w,response)




}