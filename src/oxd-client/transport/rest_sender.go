package transport

import (
	"oxd-client/constants"
	//"crypto/tls"
	"fmt"
	"strings"
	//"bytes"
	//"encoding/json"
	//"io/ioutil"
	//"net/http"
	"gopkg.in/resty.v1"
	//"net/http"
	"crypto/tls"
	//"bytes"
	//"encoding/json"
	//"io/ioutil"
	"oxd-client/model/transport"
)

type AuthSuccess struct {

}

// function which sends request via REST
func SendRest( request []byte, requestParam transport.OxdConnectionParam) []byte {

	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify:true})
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
	resty.SetDebug(requestParam.Debug)

	req:= resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(request[:])).
		//SetBody("").
			SetHeader("Accept","*/*").
		SetResult(&AuthSuccess{})

		if(requestParam.AccessToken != ""){
			req.SetHeader("Authorization", "Bearer "+ requestParam.AccessToken)
		}

		resp, _ := req.Post(getEndpointUrl(requestParam.Address,requestParam.CommandType))
		return resp.Body()
}

// Function to get url from command type
func getEndpointUrl (address string,command constants.CommandType) string {
	commandtype := fmt.Sprint(command)
	endpoint := fmt.Sprint(address+"/"+commandtype)
	return strings.Replace(endpoint, "_","-", -1)
}