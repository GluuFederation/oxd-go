package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/token"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
	"oxd-client-test/utils"
)

func TestGetAuthorizationCode(t *testing.T) {
	//BEFORE
	requestParams := model.AuthorizationCodeRequestParams{utils.RegisterClient(),
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		make([]string,0),
		"",""}
	request := client.BuildOxdRequest(constants.GET_AUTHORIZATION_CODE,requestParams)
	var response transport.OxdResponse
	var responseParams model.AuthorizationCodeResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Code,"Code should not be empty")
}
