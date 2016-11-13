package rago

import (
	"reflect"
	"strings"

	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/ralog"
)

const ACTION_OF_LIST = "List"

var methodActionMap = map[rahttp.Method]string{
	rahttp.METHOD_GET:    "Get",
	rahttp.METHOD_POST:   "Create",
	rahttp.METHOD_PUT:    "Update",
	rahttp.METHOD_DELETE: "Delete",
}

type router struct {
	request        *rahttp.Request
	response       *rahttp.Response
	uriPatterns    map[rahttp.UriPattern]rahttp.ResourceAndMethod
	resourceAction reflect.Value
}

func newRouter(request *rahttp.Request, response *rahttp.Response, uriPatterns map[rahttp.UriPattern]rahttp.ResourceAndMethod) *router {
	ralog.Debug("rago: NewRouter")

	return &router{
		request:     request,
		response:    response,
		uriPatterns: uriPatterns,
	}
}

func (r *router) match() {
	path := strings.TrimSpace(r.request.GetUriPath())

	pathSegments := strings.Split(strings.Trim(path, "/"), "/")
	ralog.Debug("rago: pathSegments ", pathSegments)
	pathSegmentCount := len(pathSegments)
	ralog.Debug("rago: pathSegmentCount ", pathSegmentCount)

	args := make(map[string]string)

	matched := false
	for pattern, resourceAndMethod := range r.uriPatterns {
		patternSegments := strings.Split(strings.Trim(string(pattern), "/"), "/")
		ralog.Debug("rago: patternSegments ", patternSegments)
		patternSegmentCount := len(patternSegments)
		ralog.Debug("rago: patternSegmentCount ", patternSegmentCount)
		if patternSegmentCount != pathSegmentCount {
			continue
		}

		matched = true
		for i := 0; i < patternSegmentCount; i++ {
			if len(patternSegments[i]) > 0 && patternSegments[i][0] == ':' {
				args[patternSegments[i][1:]] = pathSegments[i]
			} else if patternSegments[i] != pathSegments[i] {
				matched = false
				break
			}
		}

		if matched {
			r.request.WithMatchedUriPattern(pattern)
			r.request.WithAttributes(args)
			method := r.request.GetMethod()
			ralog.Debug("rago: method ", method)
			if !isMethodSupported(method, resourceAndMethod.Methods) {
				raerror.PanicWith(raerror.TypMethodNotAllowed, 0, "rago: unsupported method "+string(method))
			}

			lastSegmentIsAttribute := patternSegments[patternSegmentCount-1][0] == ':'
			ralog.Debug("rago: lastSegmentIsAttribute ", lastSegmentIsAttribute)
			r.withResourceAction(resourceAndMethod.ResourceObj, method, lastSegmentIsAttribute)
			break
		}

		if !matched {
			raerror.PanicWith(raerror.TypNotFound, 0, "rago: resource not found")
		}
	}

	ralog.Debug("rago: router.match")
}

func (r *router) withResourceAction(resourceObj interface{}, method rahttp.Method, lastSegmentIsAttribute bool) {
	resourceType := reflect.TypeOf(resourceObj)

	actionName := methodActionMap[method]
	if method == rahttp.METHOD_GET && !lastSegmentIsAttribute {
		actionName = ACTION_OF_LIST
	}
	resourcePtr := reflect.New(resourceType)
	action := resourcePtr.MethodByName(actionName)
	emtpyValue := reflect.Value{}
	if action == emtpyValue {
		raerror.PanicWith(raerror.TypRuntime, 0, "rago: resource action '"+actionName+"' not found")
	}
	ralog.Debug("rago: action ", action)

	// _, ok := resourcePtr.Type().MethodByName(actionName)
	// if !ok {
	// 	raerror.PanicWith(raerror.TypRuntime, 0, "rago: resource action '"+actionName+"' not found")
	// }

	// _, ok := resourcePtrType.MethodByName(actionName)
	// if !ok {
	// 	raerror.PanicWith(raerror.TypRuntime, 0, "rago: resource action '"+actionName+"' not found")
	// }

	// resourceType := resourcePtrType.Elem()
	// ralog.Debug("rago: resource type ", resourceType)

	// newResourcePtr := reflect.New(resourceType)
	// action := newResourcePtr.MethodByName(actionName)
	// ralog.Debug("rago: resource action ", action)
	r.resourceAction = action
}

func (r *router) callResourceAction() {
	ralog.Debug("rago: router.callResourceAction")

	r.resourceAction.Call([]reflect.Value{reflect.ValueOf(r.request), reflect.ValueOf(r.response)})
}

func isMethodSupported(method rahttp.Method, methods []rahttp.Method) bool {
	for _, item := range methods {
		if item == method {
			return true
		}
	}

	return false
}
