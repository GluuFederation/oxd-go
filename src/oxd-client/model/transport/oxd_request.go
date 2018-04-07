//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package transport

import (
	//"fmt"
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
	"strconv"
)
var COMMAND_STR_LENGTH_SIZE = 4

type Param interface {}

// OxdRequest structure which is used for all requests as a wrapper for data
type OxdRequest struct {
	Command constants.CommandType `json:"command"`
	Params Param `json:"params"`
}

// Function returns OxdRequest as a JSON byte array with length attached before. It is used for socket communication
func (r OxdRequest) ToOxdJSONWithLength() []byte{
	value , err := json.Marshal(r)
	utils.CheckError("transport.OxdRequest","JSON marshalling error",err)
	return append(getLength(value),value...)
}

// Function returns OxdRequest as a JSON byte array
func (r OxdRequest) ToOxdJSON() []byte{
	value , err := json.Marshal(r.Params)
	utils.CheckError("transport.OxdRequest","JSON marshalling error",err)
	return value
}

// Function returns size of byte array
func getLength(message []byte) []byte {
	length := strconv.Itoa(len(message))
	for len(length) < COMMAND_STR_LENGTH_SIZE {
		length = "0"+length
	}
	return []byte(length)
}
