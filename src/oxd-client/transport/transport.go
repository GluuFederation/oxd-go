//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package transport

import (
	"fmt"
	"net"
	"bytes"
	"net/http"
	"io/ioutil"
	"oxd-client/utils"
	"oxd-client/constants"
	"encoding/json"
	"crypto/tls"
	"io"
	"strconv"
	"strings"
	
)

func SocketSend( request []byte, address string) []byte {

	conn := establishConnection(address)
	_, err := conn.Write(request)
	utils.CheckError("transport.transport","Cannot write message",err)

	response := getMessage(conn, getMessageLength(conn))
	conn.Close()
	
	return response
}
func SendViaHttp( request []byte, address string,command constants.CommandType) []byte {
	var accesstoken = " "
	var requestInterface interface{}
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	requeststring := string(request[:])
	
	commandtype := fmt.Sprint(command)
	endpoint := fmt.Sprint(address+"/"+commandtype)
	endpointurl := strings.Replace(endpoint, "_","-", -1)

	findaccesstoken := bytes.Index(request, []byte("protection_access_token"))
	
	if (findaccesstoken != -1) {
			err := json.Unmarshal(request, &requestInterface)
			if err != nil {
				fmt.Println("error:", err)
			}
			requestValues := requestInterface.(map[string]interface{})
			accesstoken = requestValues["protection_access_token"].(string)
	}
	sendrequest, _ := http.NewRequest("POST", endpointurl, bytes.NewBufferString(requeststring))
	sendrequest.Header.Set("Content-Type", "application/json")
	if(accesstoken != "") {
		sendrequest.Header.Set("Authorization", "Bearer " + accesstoken)
	}
	
	client := &http.Client{Transport: tr}
    serveresponse, err := client.Do(sendrequest)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} 
	response, err := ioutil.ReadAll(serveresponse.Body)
	defer serveresponse.Body.Close()
	return response
	}

func getMessage(reader io.Reader, length int64) []byte{
	messageReader := io.LimitReader(reader,length)
	result, err := ioutil.ReadAll(messageReader)
	utils.CheckError("transport.transport","Cannot read message",err)
	
	return result
}

func getMessageLength(reader io.Reader) int64{
	lengthReader := io.LimitReader(reader,4)

	length,err :=ioutil.ReadAll(lengthReader)
	utils.CheckError("transport.transport","Cannot read message length",err)

	lengthInt,err := strconv.ParseInt(string(length),10,64)
	utils.CheckError("transport.transport","Cannot rconvert message length",err)

	return lengthInt
}

func establishConnection(address string)  *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	utils.CheckError("transport.transport","Cannot resolve ip address",err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	utils.CheckError("transport.transport","Cannot create connection",err)

	return conn
}
