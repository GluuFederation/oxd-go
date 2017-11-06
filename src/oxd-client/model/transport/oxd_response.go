//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package transport

import (
	
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
)

type OxdResponse struct {
	Status constants.Status `json:"status"`
	Params json.RawMessage `json:"data"`
}

func (r OxdResponse) GetParams(param interface{}){
	err := json.Unmarshal([]byte(r.Params), &param)
	utils.CheckError("OxdResponse","Param unmarshall error", err)

}
