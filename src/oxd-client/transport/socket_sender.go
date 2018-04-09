package transport

import (
	"oxd-client/utils"
	"net"
	"io"
	"io/ioutil"
	"strconv"
)

// function which sends request via Socket
func SendSocket( request []byte, address string) []byte {

	conn := establishConnection(address)
	_, err := conn.Write(request)
	utils.CheckError("transport.transport","Cannot write message",err)

	response := getMessage(conn, getMessageLength(conn))
	conn.Close()

	return response
}

// Establishing connection
func establishConnection(address string)  *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	utils.CheckError("transport.transport","Cannot resolve ip address",err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	utils.CheckError("transport.transport","Cannot create connection",err)

	return conn
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