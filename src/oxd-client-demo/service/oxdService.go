//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package service

import (
	
	"oxd-client/client"
	"oxd-client/model/transport"
	"encoding/json"
	"oxd-client/utils"
	"github.com/juju/loggo"
)

func CallOxdServer (request transport.OxdRequest,response *transport.OxdResponse, params transport.OxdConnectionParam){
	
	client.Send(request, params, response)
	debugCommunication(request,*response)
}

func debugCommunication(request transport.OxdRequest,response transport.OxdResponse)  {
	LOG :=loggo.GetLogger("default")
    res,err := json.Marshal(response)
	utils.CheckError("CallOxdServer","Marshal error",err)
	req, err := json.Marshal(request)
	utils.CheckError("CallOxdServer","Marshal error",err)

	LOG.Debugf("Request: "+string(req))
	LOG.Debugf("Response: "+string(res))
	
}


