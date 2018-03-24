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

func TestCheckAccessToken(t *testing.T) {
	//BEFORE
	codeResponse, oxdId := utils.ExecCodeFlow()
	requestParams := validation.CheckAccessTokenRequestParams{oxdId, codeResponse.IdToken, codeResponse.AccessToken}
	request := client.BuildOxdRequest(constants.CHECK_ACCESS_TOKEN,requestParams)
	connectionParams := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.CHECK_ACCESS_TOKEN,}
	var response transport.OxdResponse
	var responseParams validation.CheckAccessTokenResponseParams

	//TEST
	client.Send(request,connectionParams,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Active,"Active should not be empty")
	assert.NotEmpty(t,responseParams.ExpiresAt,"ExpiresAt should not be empty")
	assert.NotEmpty(t,responseParams.IssuedAt,"IssuedAt should not be empty")
}
