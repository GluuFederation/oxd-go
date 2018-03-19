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

func TestRsProtect(t *testing.T) {
	//BEFORE
	requestParams := uma.RsProtectRequestParams{utils.RegisterClient(""), []protect.RsResource{ protect.RsResource{conf.TestConfiguration.Path, conf.TestConfiguration.Condition}},""}
	requestParams.OxdId = utils.RegisterClient("")
	request := client.BuildOxdRequest(constants.RS_PROTECT,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.RS_PROTECT}

	var response transport.OxdResponse
	var responseParams uma.RsProtectResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.OxdId,"OxdId should not be empty")
}
