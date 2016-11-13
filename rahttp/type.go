package rahttp

const (
	METHOD_GET    Method = "GET"
	METHOD_POST   Method = "POST"
	METHOD_PUT    Method = "PUT"
	METHOD_DELETE Method = "DELETE"
)

type UriPattern string

type Method string

type ResourceAndMethod struct {
	ResourceObj interface{}
	Methods     []Method
}
