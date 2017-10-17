//
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

type OxdRequest struct {
	Command constants.CommandType `json:"command"`
	Params Param `json:"params"`
}

func (r OxdRequest) ToOxdJSON() []byte{
	//fmt.Println(r)
	value , err := json.Marshal(r)
	utils.CheckError("transport.OxdRequest","JSON marshalling error",err)
	return append(getLength(value),value...)
}

func getLength(message []byte) []byte {
	length := strconv.Itoa(len(message))
	for len(length) < COMMAND_STR_LENGTH_SIZE {
		length = "0"+length
	}
	return []byte(length)
}
