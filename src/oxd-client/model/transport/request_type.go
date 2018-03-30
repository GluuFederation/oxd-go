package transport

type REQUEST_TYPE string

const (
	SOCKET = REQUEST_TYPE("socket")
	REST = REQUEST_TYPE("rest")
)

func GetRequestTypeByName(name string) REQUEST_TYPE{
	switch name {
	case "SOCKET": return SOCKET
	case "REST": return REST
	}
	return REST
}