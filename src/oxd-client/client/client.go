//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package client

import (
	
	"oxd-client/model/transport"
	socket "oxd-client/transport"
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
)

func Send(request transport.OxdRequest, address string, response *transport.OxdResponse) {
	err := json.Unmarshal(socket.SocketSend(request.ToOxdJSON(), address),response)
	utils.CheckError("client.client", "Response Unmarshal error",err)
}

func BuildOxdRequest(command constants.CommandType, params transport.Param) transport.OxdRequest {
	
	return transport.OxdRequest{command,params}
}

