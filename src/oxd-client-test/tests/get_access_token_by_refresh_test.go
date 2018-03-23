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
	"oxd-client-test/utils"
	"oxd-client/model/params/token"
	"oxd-client/model/transport"
)

func TestSocketGetTokensByRefreshToken(t *testing.T) {
	executeGetTokensByRefreshTokenTest(t,utils.GetSocketRequest)
}

func TestRestGetTokensByRefreshToken(t *testing.T) {
	executeGetTokensByRefreshTokenTest(t,utils.GetRestRequest)
}

func executeGetTokensByRefreshTokenTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	refreshToken, oxdid := utils.GetRefreshToken(getRequest)
	requestParams := model.GetAccessTokenByRefreshTokenRequestParams{oxdid,
		refreshToken,"", []string{"openid"}}

	request,connectionParam := getRequest(constants.GET_ACCESS_TOKEN_BY_REFRESH_TOKEN,requestParams)
	var response transport.OxdResponse
	var responseParams model.GetAccessTokenByRefreshTokenResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AccessToken,"AccessToken should not be empty")
	assert.NotEmpty(t,responseParams.ExpiresIn,"ExpiresIn should not be empty")
	assert.NotEmpty(t,responseParams.RefreshToken,"RefreshToken should not be empty")
}
