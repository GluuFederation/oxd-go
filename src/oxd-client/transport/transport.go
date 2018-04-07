//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

// Package transport provides connection functionalities
package transport

import (
	"oxd-client/model/transport"
	"github.com/juju/loggo"
)

// Function which dispatches request depanding on request type
func Send( request transport.OxdRequest, requestParam transport.OxdConnectionParam) []byte{
	switch requestParam.Type {
	case transport.SOCKET:
		return SendSocket(request.ToOxdJSONWithLength(),requestParam.Address)
	case transport.REST:
		return SendRest(request.ToOxdJSON(),requestParam)
	default:
		loggo.GetLogger("send").Errorf("Wrong transport type")
		return nil
	}
}



