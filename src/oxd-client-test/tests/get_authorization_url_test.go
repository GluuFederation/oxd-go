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
	"oxd-client/model/params/url"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
	"oxd-client-test/utils"
)

func TestGetAuthorizationUrl(t *testing.T) {
	//BEFORE
	requestParams := model.AuthorizationUrlRequestParams{utils.RegisterClient(""),"",make([]string,0),"","",nil}
	request := client.BuildOxdRequest(constants.GET_AUTHORIZATION_URL,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET, "",constants.GET_AUTHORIZATION_URL}

	var response transport.OxdResponse
	var responseParams model.AuthorizationUrlResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AuthorizationUrl,"AuthorizationUrl should not be empty")
}
