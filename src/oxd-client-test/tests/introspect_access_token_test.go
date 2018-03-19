package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/token"
	//"oxd-client-test/conf"
	"oxd-client/model/transport"
	"oxd-client-test/utils"
)

func TestSocketIntrospectAccessToken(t *testing.T) {
	executeIntrospectAccessTokenTest(t,utils.GetSocketRequest)
}

func TestRestIntrospectAccessToken(t *testing.T) {
	executeIntrospectAccessTokenTest(t,utils.GetRestRequest)
}

func executeIntrospectAccessTokenTest(t *testing.T, getRequest utils.GetRequest){
	//BEFORE
	accessToken := utils.GetClientToken(getRequest)
	requestParams := model.IntrospectAccessTokenRequestParams{"",accessToken}

	request, connectionParam := getRequest(constants.INTROSPECT_ACCESS_TOKEN,requestParams)
	var response transport.OxdResponse
	var responseParams model.IntrospectAccessTokenResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
//	assert.NotEmpty(t,responseParams.AccessToken,"AccessToken shoudld not be empty")
//	assert.True(t,responseParams.ExpiresIn > 0,"ExpiresIn should be more than 0")
}



