//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

//Package client executes calls and builds requests
package client

import (
	"oxd-client/model/transport"
	sender "oxd-client/transport"
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
)

//Function executes call to Oxd Server
func Send(request transport.OxdRequest, requestParam transport.OxdConnectionParam, response *transport.OxdResponse) {
	error := json.Unmarshal(sender.Send(request, requestParam),response)
	utils.CheckError("client.client", "Response Unmarshal error",error)
}

//Oxd Request Builder
func BuildOxdRequest(command constants.CommandType, params transport.Param) transport.OxdRequest {
	return transport.OxdRequest{command,params}
}

