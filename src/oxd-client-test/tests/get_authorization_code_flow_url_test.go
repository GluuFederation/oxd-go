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
	"oxd-client/model/params/url"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
	"oxd-client-test/utils"
)

func TestGetAuthorizationCodeFlowUrl(t *testing.T) {
	//BEFORE
	requestParams := model.AuthorizationCodeFlowRequestParams{
		utils.RegisterClient(""),
		conf.TestConfiguration.RedirectUrl,
		conf.TestConfiguration.ClientId,
		conf.TestConfiguration.ClientSecret,
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		conf.TestConfiguration.Scope,"",""}
	request := client.BuildOxdRequest(constants.AUTHORIZATION_CODE_FLOW,requestParams)
	connectionParam := transport.OxdConnectionParam{conf.TestConfiguration.Host,transport.SOCKET,"",constants.CHECK_ID_TOKEN}

	var response transport.OxdResponse
	var responseParams model.AuthorizationCodeFlowResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.AuthorizationCode,"AuthorizationCode should not be empty")
	assert.NotEmpty(t,responseParams.AccessToken,"AccessToken should not be empty")
	assert.NotEmpty(t,responseParams.ExpiresInSeconds,"ExpiresInSeconds should not be empty")
	assert.NotEmpty(t,responseParams.Scope,"Scope should not be empty")
	assert.NotEmpty(t,responseParams.IdToken,"IdToken should not be empty")
	assert.NotEmpty(t,responseParams.RefreshToken,"RefreshToken should not be empty")
}

