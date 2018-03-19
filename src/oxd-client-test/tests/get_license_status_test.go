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
	"oxd-client-test/conf"
)

func TestLicenseStatus(t *testing.T) {
	//BEFORE
	requestParams := validation.LicenseStatusRequestParams{}
	request := client.BuildOxdRequest(constants.LICENSE_STATUS,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.LICENSE_STATUS}

	var response transport.OxdResponse
	var responseParams validation.LicenseStatusResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
}
