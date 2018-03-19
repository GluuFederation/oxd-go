//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 03/01/17
//
package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/uma"
	"oxd-client/model/params/uma/protect"
	"oxd-client-test/utils"
	"oxd-client-test/conf"
	"oxd-client/model/transport"
)

func TestUma(t *testing.T) {
	oxdId := utils.RegisterClient("")
	protectRs(oxdId)
	checkAccessParams := checkAccess(oxdId,"", "/ws/phone", "GET")

	assert.Equal(t,"denied",checkAccessParams.Access, "Access should be denied")
	assert.NotEmpty(t,checkAccessParams.Ticket,"","Ticket should not be empty")
	ticket := checkAccessParams.Ticket

	rpt := obtainRpt(oxdId)
	rptCheckAccessParams := checkAccess(oxdId,rpt,"/ws/phone", "GET")
	assert.Equal(t,"denied",rptCheckAccessParams.Access, "Access should be denied")

	authorizeRpt(oxdId,rpt,ticket)
	rptAuthCheckAccessParams := checkAccess(oxdId,rpt,"/ws/phone", "GET")
	assert.Equal(t,"granted",rptAuthCheckAccessParams.Access, "Access should be denied")

	assertNotProtectedError(oxdId,rpt,t)
}

func protectRs(oxdID string) {
	requestParams := uma.RsProtectRequestParams{oxdID,
		[]protect.RsResource{ protect.RsResource{conf.TestConfiguration.Path, conf.TestConfiguration.Condition}},""}
	request := client.BuildOxdRequest(constants.RS_PROTECT,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RS_PROTECT}
	var response transport.OxdResponse

	client.Send(request,connectionParam,&response)
}

func checkAccess(oxdID string, rpt string, path string, httpMethod string ) uma.RsCheckAccessResponseParams{
	requestParams := uma.RsCheckAccessRequestParams{oxdID,rpt,path,httpMethod,""}
	request := client.BuildOxdRequest(constants.RS_CHECK_ACCESS,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RS_CHECK_ACCESS}

	var response transport.OxdResponse
	var responseParams uma.RsCheckAccessResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams
}

func obtainRpt(oxdID string) string {
	requestParams := uma.RpGetRptRequestParams{oxdID,"","","","","","","",""}
	request := client.BuildOxdRequest(constants.RP_GET_RPT,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RP_GET_RPT}

	var response transport.OxdResponse
	var responseParams uma.RpGetRptResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams.Rpt
}

func authorizeRpt(oxdID string, rpt string, ticket string){
	requestParams := uma.RpAuthorizeRptRequestParams{oxdID,rpt,ticket}
	request := client.BuildOxdRequest(constants.RP_AUTHORIZE_RPT,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RP_AUTHORIZE_RPT}

	var response transport.OxdResponse
	var responseParams uma.RpAuthorizeRptResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
}

func assertNotProtectedError(oxdID string, rpt string, t *testing.T){
	requestParams := uma.RsCheckAccessRequestParams{oxdID,rpt,"/no/such/path","GET",""}
	request := client.BuildOxdRequest(constants.RS_CHECK_ACCESS,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RS_CHECK_ACCESS}

	var response transport.OxdResponse

	client.Send(request,connectionParam,&response)

	assert.Equal(t,constants.STATUS_ERROR,response.Status)
}