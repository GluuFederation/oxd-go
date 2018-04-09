package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/token"
	"oxd-client-test/conf"
	"oxd-client/model/transport"
	"oxd-client-test/utils"
)

func TestSocketGetClientToken(t *testing.T) {
	executeGetClientTokenTest(t,utils.GetSocketRequest)
}

func TestRestGetClientToken(t *testing.T) {
	executeGetClientTokenTest(t,utils.GetRestRequest)
}

func executeGetClientTokenTest(t *testing.T, getRequest utils.GetRequest) {
	//BEFORE
	requestParams := model.GetClientTokenRequestParams {"",
		conf.TestConfiguration.ClientId,
		conf.TestConfiguration.ClientSecret,
		conf.TestConfiguration.OpHost,
		"",[]string{"uma_protection"}}

	request, connectionParam := getRequest(constants.GET_CLIENT_TOKEN,requestParams)
	var response transport.OxdResponse
	var responseParams model.GetClientTokenResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AccessToken,"AccessToken shoudld not be empty")
	assert.True(t,responseParams.ExpiresIn > 0,"ExpiresIn should be more than 0")
}



