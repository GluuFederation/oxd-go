package transport

// Request type
type REQUEST_TYPE string

const (
	SOCKET = REQUEST_TYPE("socket")
	REST = REQUEST_TYPE("rest")
)

// Function returns type of request by name
func GetRequestTypeByName(name string) REQUEST_TYPE{
	switch name {
	case "SOCKET": return SOCKET
	case "REST": return REST
	}
	return REST
}