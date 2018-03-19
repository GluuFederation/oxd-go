//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package transport

import (
	"oxd-client/model/transport"
)


func Send( request transport.OxdRequest, requestParam transport.OxdConnectionParam) []byte{
	switch requestParam.Type {
	case transport.SOCKET:
		return SendSocket(request.ToOxdJSONWithLength(),requestParam.Address)
	case transport.REST:
		return SendRest(request.ToOxdJSON(),requestParam)
	}
	//TODO error handling
	return nil
}



