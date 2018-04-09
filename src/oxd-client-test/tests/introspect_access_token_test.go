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
	oxdId := utils.RegisterClientSite(getRequest)
	accessToken := utils.GetClientTokenByOxdId(getRequest, oxdId)
	requestParams := model.IntrospectAccessTokenRequestParams{oxdId,accessToken}

	request, connectionParam := getRequest(constants.INTROSPECT_ACCESS_TOKEN,requestParams)
	var response transport.OxdResponse
	var responseParams model.IntrospectAccessTokenResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
}



