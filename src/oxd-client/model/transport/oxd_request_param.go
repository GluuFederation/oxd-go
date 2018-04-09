//  Author: Michał Kępkowski
//  Date: 07.04.2018

package transport

import "oxd-client/constants"

// Oxd connection parameters which specifies address, type(REST/Socket), access token and command type
type OxdConnectionParam struct {
	Address string
	Type REQUEST_TYPE
	AccessToken string
	CommandType constants.CommandType
	Debug bool
}