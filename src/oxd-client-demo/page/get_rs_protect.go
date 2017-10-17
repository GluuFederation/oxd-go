package page

import (
	
	
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/uma"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client/model/params/uma/protect"
	"oxd-client-demo/service"
	"oxd-client-demo/utils"
	
)

func ProtectResourcesPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars ,  accesstoken string,globalvariables conf.GlobalVars) {

   var oxdResponse transport.OxdResponse
   requestParams := uma.RsProtectRequestParams{globalvariables.Oxdid, []protect.RsResource{ protect.RsResource{configuration.Path, configuration.Condition}}, accesstoken}
    
   page.CallOxdServer(
	client.BuildOxdRequest(constants.RS_PROTECT,requestParams),
	&oxdResponse,
	globalvariables.Host)

	var response uma.RsProtectResponseParams
	
	oxdResponse.GetParams(&response)
	utils.DisplayResponse(w,response)
	



}
  
	









