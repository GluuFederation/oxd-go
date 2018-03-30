package page

import (
	"oxd-client/constants"
	"oxd-client/client"
	"oxd-client/model/params/registration"
	"oxd-client-demo/conf"
	"oxd-client-demo/utils"
	"oxd-client/model/transport"
)


func RegisterClientSite(conf *conf.Configuration) string{
	requestParams := prepareRegisterParams(*conf)

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.REGISTER_SITE, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.RegisterSiteParams.OxdId = responseParams.OxdId
	return responseParams.OxdId;
}

func prepareRegisterParams(conf conf.Configuration) model.RegisterSiteRequestParams{
	var requestParams model.RegisterSiteRequestParams
	requestParams.OpHost = conf.IdpHost
	requestParams.AuthorizationRedirectUri = conf.RegisterSiteParams.AuthorizationRedirectUri
	requestParams.PostLogoutRedirectUri = conf.RegisterSiteParams.PostLogoutRedirectUri
	requestParams.ClientLogoutUri = conf.RegisterSiteParams.ClientLogoutUri
	requestParams.RedirectUris = conf.RegisterSiteParams.RedirectUris
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = conf.RegisterSiteParams.Scope
	requestParams.GrantTypes = conf.RegisterSiteParams.GrantTypes
	requestParams.ResponseTypes = conf.RegisterSiteParams.ResponseTypes
	requestParams.ClientId = conf.ClientId
	requestParams.ClientSecret = conf.ClientSecret
	requestParams.ClientRegistrationClientUri = conf.ClientRegistrationClientUri
	return  requestParams
}
//
////----------------------UPDATE-----------------------------------
//
//func TestSocketUpdateClientSite(t *testing.T) {
//	executeUpdateClientSiteTest(t,utils.GetSocketRequest)
//}
//
//func TestRestUpdateClientSite(t *testing.T) {
//	executeUpdateClientSiteTest(t,utils.GetRestRequest)
//}
//
//func executeUpdateClientSiteTest(t *testing.T,getRequest utils.GetRequest) {
//	//BEFORE
//	updateParams := utils.PrepareUpdateParams()
//	updateRequest, connectionParam := getRequest(constants.UPDATE_SITE,updateParams)
//
//	var updateResponse transport.OxdResponse
//	var updateResponseParams model.UpdateSiteResponseParams
//
//	//TEST
//	client.Send(updateRequest,connectionParam,&updateResponse)
//
//	//ASSERT
//	updateResponse.GetParams(&updateResponseParams)
//	assert.Equal(t,constants.STATUS_OK,updateResponse.Status,"Status should be ok")
//	assert.NotEmpty(t,updateResponseParams.OxdId,"OxdId should not be empty")
//}
//
////----------------------REMOVE-----------------------------------
//
//func TestSocketRemoveClientSite(t *testing.T) {
//	executeRemoveClientSiteTest(t,utils.GetSocketRequest)
//}
//
//func TestRestRemoveClientSite(t *testing.T) {
//	executeRemoveClientSiteTest(t,utils.GetRestRequest)
//}
//
//func executeRemoveClientSiteTest(t *testing.T,getRequest utils.GetRequest) {
//	//BEFORE
//	removeParams := model.RemoveSiteRequestParams{utils.RegisterClientSite(getRequest)}
//	removeRequest, connectionParam := getRequest(constants.REMOVE_SITE, removeParams)
//
//	var removeResponse transport.OxdResponse
//	var removeResponseParams model.RemoveSiteResponseParams
//
//	//TEST
//	client.Send(removeRequest,connectionParam,&removeResponse)
//
//	//ASSERT
//	removeResponse.GetParams(&removeResponseParams)
//	assert.Equal(t,constants.STATUS_OK, removeResponse.Status,"Status should be ok")
//	assert.NotEmpty(t, removeResponseParams.OxdId,"OxdId should not be empty")
//}
