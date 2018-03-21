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
	//"oxd-client-test/conf"
	"oxd-client-test/utils"
)

func TestSocketGetAuthorizationUrl(t *testing.T) {
	executeGetAuthorizationUrlTest(t,utils.GetSocketRequest)
}

func TestRestGetAuthorizationUrl(t *testing.T) {
	executeGetAuthorizationUrlTest(t,utils.GetRestRequest)
}

func executeGetAuthorizationUrlTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	requestParams := utils.PrepareAuthorizationUrlRequestParams(getRequest)
	request, connectionParam := getRequest(constants.GET_AUTHORIZATION_URL, requestParams)

	var response transport.OxdResponse
	var responseParams model.AuthorizationUrlResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AuthorizationUrl,"AuthorizationUrl should not be empty")
}

