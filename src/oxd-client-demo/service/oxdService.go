package page

import (
	"oxd-client/client"
	"oxd-client/model/transport"
	"encoding/json"
	"oxd-client/utils"
	"github.com/juju/loggo"
)

func CallOxdServer (request transport.OxdRequest,response *transport.OxdResponse, host string){
	client.Send(request, host, response)
	debugCommunication(request,*response)
}

func debugCommunication(request transport.OxdRequest,response transport.OxdResponse){
	LOG :=loggo.GetLogger("CallOxdServer")

	res,err := json.Marshal(response)
	utils.CheckError("CallOxdServer","Marshal error",err)
	req, err := json.Marshal(request)
	utils.CheckError("CallOxdServer","Marshal error",err)

	LOG.Debugf("Request: "+string(req))
	LOG.Debugf("Response: "+string(res))
}


