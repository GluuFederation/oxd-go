package transport

type REQUEST_TYPE string

const (
	SOCKET = REQUEST_TYPE("socket")
	REST = REQUEST_TYPE("rest")
)