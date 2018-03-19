//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package tests

import (
	"testing"
	"oxd-client/model/transport"
	"oxd-client/model/params/registration"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"

	"oxd-client-test/utils"
	//"oxd-client-test/conf"
)


//------------------REGISTER----------------------------------
func TestSocketRegisterClientSite(t *testing.T) {
	executeRegisterClientSiteTest(t,utils.GetSocketRequest)
}

func TestRestRegisterClientSite(t *testing.T) {
	executeRegisterClientSiteTest(t,utils.GetRestRequest)
}

func executeRegisterClientSiteTest(t *testing.T, getRequest utils.GetRequest) {
	//BEFORE
	requestParams := utils.PrepareRegisterParams(utils.GetClientToken(getRequest))
	request, connectionParam := getRequest(constants.REGISTER_SITE,requestParams)

	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.OxdId,"OxdId should not be empty")
}

//----------------------UPDATE-----------------------------------

func TestSocketUpdateClientSite(t *testing.T) {
	executeUpdateClientSiteTest(t,utils.GetSocketRequest)
}

func TestRestUpdateClientSite(t *testing.T) {
	executeUpdateClientSiteTest(t,utils.GetRestRequest)
}

func executeUpdateClientSiteTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	updateParams := utils.PrepareUpdateParams()
	updateRequest, connectionParam := getRequest(constants.UPDATE_SITE,updateParams)

	var updateResponse transport.OxdResponse
	var updateResponseParams model.UpdateSiteResponseParams

	//TEST
	client.Send(updateRequest,connectionParam,&updateResponse)

	//ASSERT
	updateResponse.GetParams(&updateResponseParams)
	assert.Equal(t,constants.STATUS_OK,updateResponse.Status,"Status should be ok")
	assert.NotEmpty(t,updateResponseParams.OxdId,"OxdId should not be empty")
}

//----------------------REMOVE-----------------------------------

func TestSocketRemoveClientSite(t *testing.T) {
	executeRemoveClientSiteTest(t,utils.GetSocketRequest)
}

func TestRestRemoveClientSite(t *testing.T) {
	executeRemoveClientSiteTest(t,utils.GetRestRequest)
}

func executeRemoveClientSiteTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	removeParams := model.RemoveSiteRequestParams{utils.RegisterClientSite(getRequest)}
	removeRequest, connectionParam := getRequest(constants.REMOVE_SITE, removeParams)

	var removeResponse transport.OxdResponse
	var removeResponseParams model.RemoveSiteResponseParams

	//TEST
	client.Send(removeRequest,connectionParam,&removeResponse)

	//ASSERT
	removeResponse.GetParams(&removeResponseParams)
	assert.Equal(t,constants.STATUS_OK, removeResponse.Status,"Status should be ok")
	assert.NotEmpty(t, removeResponseParams.OxdId,"OxdId should not be empty")
}

