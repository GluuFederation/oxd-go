package transport

import "oxd-client/constants"

type OxdConnectionParam struct {
	Address string
	Type REQUEST_TYPE
	AccessToken string
	CommandType constants.CommandType
}