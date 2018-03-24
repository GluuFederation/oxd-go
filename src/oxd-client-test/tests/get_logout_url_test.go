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


func TestSocketGetLogoutUrl(t *testing.T) {
	executeGetLogoutUrlTest(t,utils.GetSocketRequest)
}

func TestRestGetLogoutUrl(t *testing.T) {
	executeGetLogoutUrlTest(t,utils.GetRestRequest)
}

func executeGetLogoutUrlTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	tokensResponse, oxdId := utils.GetTokens(getRequest)
	requestParams := model.LogoutUrlRequestParams{oxdId,"",tokensResponse.IdToken,
		conf.TestConfiguration.PostLogoutRedirectUrl,"",""}

	request,connectionParam := getRequest(constants.GET_LOGOUT_URI,requestParams)
	var response transport.OxdResponse
	var responseParams model.LogoutUrlResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Uri,"Uri should not be empty")
}
