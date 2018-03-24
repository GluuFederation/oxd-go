//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 03/01/17
//
package tests

import (
	"testing"
	"oxd-client/model/params/uma"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client-test/utils"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
	"oxd-client/model/params/uma/protect"
)

func TestSocketUMARsProtect(t *testing.T) {
	executeUMARsProtectTest(t,utils.GetSocketRequest)
}

func TestRestUMARsProtect(t *testing.T) {
	executeUMARsProtectTest(t,utils.GetRestRequest)
}

func executeUMARsProtectTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	oxdId := utils.RegisterClientSite(getRequest)
	requestParams := uma.RsProtectRequestParams{oxdId,
	[]protect.RsResource{ protect.RsResource{conf.TestConfiguration.Path, conf.TestConfiguration.Condition}},
	""}

	request,connectionParam := getRequest(constants.RS_PROTECT,requestParams)
	var response transport.OxdResponse
	var responseParams uma.RsProtectResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.OxdId,"OxdId should not be empty")
}