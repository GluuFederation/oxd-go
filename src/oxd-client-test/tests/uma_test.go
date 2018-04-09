//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 03/01/17
//
package tests

import (
	"github.com/stretchr/testify/assert"
	//"oxd-client-test/conf"
	"oxd-client-test/utils"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/params/uma"
	"oxd-client/model/transport"
	"testing"
)

func TestSocketUMA(t *testing.T) {
	executeUMATest(t, utils.GetSocketRequest)
}

func TestRestUMA(t *testing.T) {
	executeUMATest(t, utils.GetRestRequest)
}

func executeUMATest(t *testing.T, getRequest utils.GetRequest) {
	setupParams := utils.SetupClient()
	utils.ProtectRs(setupParams.OxdId, getRequest)
	checkAccessParams := utils.CheckAccess(setupParams.OxdId, "", "/ws/phone", "GET", getRequest)

	assert.Equal(t, "denied", checkAccessParams.Access, "Access should be denied")
	assert.NotEmpty(t, checkAccessParams.Ticket, "", "Ticket should not be empty")
	ticket := checkAccessParams.Ticket

	rpt := obtainRpt(setupParams.OxdId, ticket, getRequest)
	rptCheckAccessParams := utils.CheckAccess(setupParams.OxdId, rpt, "/ws/phone", "GET", getRequest)
	assert.Equal(t, "granted", rptCheckAccessParams.Access, "Access should be denied")

	introspectResponse := introspectRpt(setupParams.OxdId, rpt, getRequest)
	assert.Equal(t, introspectResponse.Active, true)

	notProtectedResponse := notProtectedError(setupParams.OxdId, rpt, getRequest)
	assert.Equal(t, constants.STATUS_ERROR, notProtectedResponse.Status)
}

func introspectRpt(oxdID string, rpt string, getRequest utils.GetRequest) uma.IntrospectRptResponseParams {
	requestParams := uma.IntrospectRptRequestParams{oxdID, rpt}
	request, connectionParam := getRequest(constants.INTROSPECT_RPT, requestParams)

	var response transport.OxdResponse
	var responseParams uma.IntrospectRptResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	return responseParams
}

func obtainRpt(oxdID string, ticket string, getRequest utils.GetRequest) string {
	requestParams := uma.RpGetRptRequestParams{oxdID, ticket, "", "", "", "", []string{}, "", ""}
	request, connectionParam := getRequest(constants.RP_GET_RPT, requestParams)

	var response transport.OxdResponse
	var responseParams uma.RpGetRptResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	return responseParams.Rpt
}

func notProtectedError(oxdID string, rpt string, getRequest utils.GetRequest) transport.OxdResponse {
	requestParams := uma.RsCheckAccessRequestParams{oxdID, rpt, "/no/such/path", "GET", ""}
	request, connectionParam := getRequest(constants.RS_CHECK_ACCESS, requestParams)

	var response transport.OxdResponse

	client.Send(request, connectionParam, &response)
	return response
}
