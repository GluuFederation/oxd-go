package page

import (
	"oxd-client/constants"
	"oxd-client/client"
	"oxd-client/model/params/uma"
	"oxd-client-demo/conf"
	"oxd-client/model/params/uma/protect"
	"oxd-client/model/transport"
	"oxd-client-demo/utils"
)

func UmaRsProtect(conf *conf.Configuration) string  {
	requestParams := uma.RsProtectRequestParams{
		conf.RegisterSiteParams.OxdId,
		prepareResources(*conf),
		"",
	}

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.RS_PROTECT, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams uma.RsProtectResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.Uma.ResourcesOxdId = responseParams.OxdId
	return responseParams.OxdId
}

func prepareResources(conf conf.Configuration) []protect.RsResource{
	resources := []protect.RsResource{}
	for _,resource := range conf.Uma.Resources{
		conditions := []protect.Condition{}
		for _,condition := range resource.Conditions{
			conditions = append(conditions, protect.Condition{condition.HttpMethods, condition.Scopes, condition.TicketScopes,nil})
		}
		resources = append(resources, protect.RsResource{resource.Path,conditions})
	}
	return resources
}

func UmaCheckAccess(conf *conf.Configuration) string  {
	requestParams := uma.RsCheckAccessRequestParams{
		conf.RegisterSiteParams.OxdId,
		conf.Uma.Rpt,
		conf.Uma.CheckAccessPath,
		conf.Uma.CheckAccessMethod,
		"",
	}

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.RS_CHECK_ACCESS, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams uma.RsCheckAccessResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.Uma.Ticket = responseParams.Ticket
	return responseParams.Access
}

func UmaGetRpt(conf *conf.Configuration) string {
	requestParams := uma.RpGetRptRequestParams{}
	requestParams.OxdId = conf.RegisterSiteParams.OxdId
	requestParams.Ticket = conf.Uma.Ticket
	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.RP_GET_RPT, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams uma.RpGetRptResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.Uma.Rpt = responseParams.Rpt
	return responseParams.Rpt
}

func UmaGetClaimsGathering(conf *conf.Configuration) string {

	requestParams := uma.RpGetClaimsGatheringUrlRequestParams{conf.RegisterSiteParams.OxdId, conf.Uma.Ticket,
		conf.RegisterSiteParams.ClaimsRedirectUri[0], ""}

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.RP_GET_CLAIMS_GATHERING_URL, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams uma.RpGetClaimsGatheringUrlResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.Uma.ClaimsUrl = responseParams.URL
	return responseParams.URL
}

func UmaIntrospectRpt (conf *conf.Configuration) interface{} {
	requestParams := uma.IntrospectRptRequestParams{conf.RegisterSiteParams.OxdId, conf.Uma.Rpt}

	requestType := transport.GetRequestTypeByName(conf.Type)
	address := utils.GetAddressForType(*conf, requestType)
	request, connectionParam := utils.GetRequest(address, requestType, constants.INTROSPECT_RPT, requestParams)
	connectionParam.AccessToken = conf.AccessToken

	var response transport.OxdResponse
	var responseParams uma.IntrospectRptResponseParams

	client.Send(request, connectionParam, &response)

	response.GetParams(&responseParams)
	conf.Uma.IntrospectRpt = responseParams
	return responseParams
}