package utils

import (
	"oxd-client/model/params/uma"
	"oxd-client/constants"
	"oxd-client/client"
	"oxd-client/model/transport"
	"oxd-client/model/params/uma/protect"
	"oxd-client-test/conf"
)

func CheckAccess(oxdID string, rpt string, path string, httpMethod string , getRequest GetRequest) uma.RsCheckAccessResponseParams{
	requestParams := uma.RsCheckAccessRequestParams{oxdID,rpt,path,httpMethod,""}
	request,connectionParam := getRequest(constants.RS_CHECK_ACCESS,requestParams)

	var response transport.OxdResponse
	var responseParams uma.RsCheckAccessResponseParams

	client.Send(request,connectionParam,&response)

	response.GetParams(&responseParams)
	return responseParams
}

func ProtectRs(oxdId string, getRequest GetRequest) {
	requestParams := uma.RsProtectRequestParams{oxdId,
		[]protect.RsResource{ protect.RsResource{conf.TestConfiguration.Path,
			conf.TestConfiguration.Condition}},
		""}

	request,connectionParam := getRequest(constants.RS_PROTECT,requestParams)
	var response transport.OxdResponse
	var responseParams uma.RsProtectResponseParams

	client.Send(request,connectionParam,&response)
	response.GetParams(&responseParams)
}