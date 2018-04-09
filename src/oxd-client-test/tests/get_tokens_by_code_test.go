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
	//"oxd-client-test/conf"
)

func TestSocketGetTokensByCode(t *testing.T) {
	executeGetTokensByCodeTest(t,utils.GetSocketRequest)
}

func TestRestGetTokensByCode(t *testing.T) {
	executeGetTokensByCodeTest(t,utils.GetRestRequest)
}

func executeGetTokensByCodeTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	codeResponses := utils.GetCode(getRequest)
	requestParams := model.TokensByCodeRequestParams{codeResponses.Get("oxdId"),
	"",codeResponses.Get("code"), codeResponses.Get("state")}

	request,connectionParam := getRequest(constants.GET_TOKENS_BY_CODE,requestParams)
	var response transport.OxdResponse
	var responseParams model.TokensByCodeResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AccessToken,"AccessToken should not be empty")
	assert.NotEmpty(t,responseParams.ExpiresIn,"ExpiresIn should not be empty")
	assert.NotEmpty(t,responseParams.IdToken,"IdToken should not be empty")
}
