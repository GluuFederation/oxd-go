//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package client

import (
	"net/url"
	"oxd-client/model/transport"
	socket "oxd-client/transport"
	https "oxd-client/transport"
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
)

func Send(request transport.OxdRequest, address string, response *transport.OxdResponse) {
	
	checkaddress, err := url.Parse(address)
	println(checkaddress)
    if err != nil {	
	 error := json.Unmarshal(socket.SocketSend(request.ToOxdJSON(), address),response)
	 utils.CheckError("client.client", "Response Unmarshal error",error)
	} else {
	 httpserror := json.Unmarshal(https.SendViaHttp(request.ToOxdJSONHttps(), address, request.Command ),response)
	 utils.CheckError("client.client", "Response Unmarshal error",httpserror)
	}
	
	
}

func BuildOxdRequest(command constants.CommandType, params transport.Param) transport.OxdRequest {
	
	return transport.OxdRequest{command,params}
}

