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
	)

func TestSocketRpGetRpt(t *testing.T) {
	executeRpGetRptTest(t,utils.GetSocketRequest)
}

func TestRestRpGetRpt(t *testing.T) {
	executeRpGetRptTest(t,utils.GetRestRequest)
}

func executeRpGetRptTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	setupParams := utils.SetupClient()
	utils.ProtectRs(setupParams.OxdId, getRequest)
	checkAccessParams := utils.CheckAccess(setupParams.OxdId,"", "/ws/phone", "GET", getRequest)
	requestParams := uma.RpGetRptRequestParams{}
	requestParams.OxdId = setupParams.OxdId
	requestParams.Ticket = checkAccessParams.Ticket
	request,connectionParam := getRequest(constants.RP_GET_RPT,requestParams)

	var response transport.OxdResponse
	var responseParams uma.RpGetRptResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Rpt,"Rpt should not be empty")
}