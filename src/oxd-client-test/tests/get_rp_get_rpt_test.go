//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client-test/utils"
	"oxd-client/model/params/uma"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
)

func TestRpGetRpt(t *testing.T) {
	//BEFORE
	requestParams := uma.RpGetRptRequestParams{}
	requestParams.OxdId = utils.RegisterClient("")
	request := client.BuildOxdRequest(constants.RP_GET_RPT,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RP_GET_RPT}

	var response transport.OxdResponse
	var responseParams uma.RpGetRptResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Rpt,"Rpt should not be empty")
}