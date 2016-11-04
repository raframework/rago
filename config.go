package rago

type UriPattern string

type Method string

type ResourceMethod struct {
	Resource interface{}
	Methods  []Method
}
