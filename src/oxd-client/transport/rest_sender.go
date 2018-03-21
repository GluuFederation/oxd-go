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

func SendRest( request []byte, requestParam transport.OxdConnectionParam) []byte {

	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify:true})
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
	resty.SetDebug(true)

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

//func SendRest( request []byte, address string,command constants.CommandType) []byte {
//	var accesstoken = " "
//	var requestInterface interface{}
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//		//, CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_256_CBC_SHA}
//		//},
//	}
//
//	requeststring := string(request[:])
//
//	commandtype := fmt.Sprint(command)
//	endpoint := fmt.Sprint(address+"/"+commandtype)
//	endpointurl := strings.Replace(endpoint, "_","-", -1)
//
//	findaccesstoken := bytes.Index(request, []byte("protection_access_token"))
//
//	if (findaccesstoken != -1) {
//		err := json.Unmarshal(request, &requestInterface)
//		if err != nil {
//			fmt.Println("error:", err)
//		}
//		requestValues := requestInterface.(map[string]interface{})
//		accesstoken = requestValues["protection_access_token"].(string)
//	}
//	sendrequest, _ := http.NewRequest("POST", endpointurl, bytes.NewBufferString(requeststring))
//	sendrequest.Header.Set("Content-Type", "application/json")
//	if(accesstoken != "") {
//		sendrequest.Header.Set("Authorization", "Bearer " + accesstoken)
//	}
//	sendrequest.Close = true
//	sendrequest.Header.Set("Accept-Encoding", "identity")
//
//	client := &http.Client{Transport: tr}
//	serveresponse, err := client.Do(sendrequest)
//	if err != nil {
//		fmt.Printf("The HTTP request failed with error %s\n", err)
//	}
//	response, err := ioutil.ReadAll(serveresponse.Body)
//	defer serveresponse.Body.Close()
//	return response
//}

func getEndpointUrl (address string,command constants.CommandType) string {
	commandtype := fmt.Sprint(command)
	endpoint := fmt.Sprint(address+"/"+commandtype)
	return strings.Replace(endpoint, "_","-", -1)
}