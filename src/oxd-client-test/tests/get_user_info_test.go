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
	"oxd-client/model/params"
	"oxd-client-test/utils"
	"oxd-client/model/transport"
)

func TestSocketGetUserInfo(t *testing.T) {
	executeGetUserInfoTest(t,utils.GetSocketRequest)
}

func TestRestGetUserInfo(t *testing.T) {
	executeGetUserInfoTest(t,utils.GetRestRequest)
}

func executeGetUserInfoTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	tokenResponse, oxdId := utils.GetTokens(getRequest)
	requestParams := model.UserInfoRequestParams{oxdId,"",
		tokenResponse.AccessToken}

	request,connectionParam := getRequest(constants.GET_USER_INFO,requestParams)
	var response transport.OxdResponse
	var responseParams model.UserInfoResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Claims,"Claims should not be empty")
}