package tests

import (
"testing"
"oxd-client/client"
"oxd-client/constants"
"github.com/stretchr/testify/assert"
//"oxd-client-test/conf"
"oxd-client/model/transport"
"oxd-client-test/utils"
	"oxd-client/model/params/registration"
)

func TestSocketSetupClient(t *testing.T) {
	executeSetupClientTest(t,utils.GetSocketRequest)
}

func TestRestSetupClient(t *testing.T) {
	executeSetupClientTest(t,utils.GetRestRequest)
}

func executeSetupClientTest(t *testing.T, getRequest utils.GetRequest){
	//BEFORE
	requestParams := utils.PrepareSetupParams()

	request, connectionParam := getRequest(constants.SETUP_CLIENT,requestParams)
	var response transport.OxdResponse
	var responseParams model.SetupClientResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
}



