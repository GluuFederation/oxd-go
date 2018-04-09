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
	"oxd-client/model/params/validation"
	"oxd-client/model/transport"
	"oxd-client-test/utils"
)

func TestSocketLicenseStatus(t *testing.T) {
	executeLicenseStatusTest(t,utils.GetSocketRequest)
}

//func TestRestLicenseStatus(t *testing.T) {
//	executeLicenseStatusTest(t,utils.GetRestRequest)
//}

func executeLicenseStatusTest(t *testing.T, getRequest utils.GetRequest) {
	//BEFORE
	requestParams := validation.LicenseStatusRequestParams{}
	request, connectionParam := getRequest(constants.LICENSE_STATUS, requestParams)

	var response transport.OxdResponse
	var responseParams validation.LicenseStatusResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
}
