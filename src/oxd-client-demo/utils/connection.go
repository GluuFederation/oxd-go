package utils

import (
	"oxd-client/model/transport"
	"oxd-client-demo/conf"
	"oxd-client/constants"
	"oxd-client/client"
)

func GetAddressForType(configuration conf.Configuration, requestType transport.REQUEST_TYPE) string {
	switch requestType {
	case transport.SOCKET: return configuration.Host
	case transport.REST: return configuration.OxdHttpsHost
	}
	return ""
	}

func GetRequest (address string, requestType transport.REQUEST_TYPE, command constants.CommandType, params transport.Param) (transport.OxdRequest, transport.OxdConnectionParam){
	request := client.BuildOxdRequest(command,params)
	connectionParam := transport.OxdConnectionParam {
		address,
		requestType,
		"",
		command}
	return request, connectionParam
}

//func GetRestRequest (command constants.CommandType, params transport.Param,  address string) (transport.OxdRequest, transport.OxdConnectionParam){
//	var accesstoken = ""
//	request := client.BuildOxdRequest(command,params)
//	connectionParam := transport.OxdConnectionParam {address,
//		transport.REST,
//		accesstoken,
//		command}
//
//	return request, connectionParam
//}