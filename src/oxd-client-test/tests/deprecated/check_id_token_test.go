//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package deprecated

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client-test/utils"
	"oxd-client/model/params/validation"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
)

func TestCheckIdToken(t *testing.T) {
	//BEFORE
	codeResponse, oxdId := utils.ExecCodeFlow()
	requestParams := validation.CheckIdTokenRequestParams{oxdId, codeResponse.IdToken}
	request := client.BuildOxdRequest(constants.CHECK_ID_TOKEN,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.CHECK_ID_TOKEN}

	var response transport.OxdResponse
	var responseParams validation.CheckIdTokenResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.ExpiresAt,"ExpiresAt should not be empty")
	assert.NotEmpty(t,responseParams.IssuedAt,"IssuedAt should not be empty")
	assert.NotEmpty(t,responseParams.Claims,"AccessToken should not be empty")
}
