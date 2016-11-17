package rahttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/ralog"
)

type Request struct {
	stdRequest        *http.Request
	matchedUriPattern UriPattern
	attributes        map[string]string
	bodyParsed        map[string]interface{}
	queryParams       map[string]interface{}
}

func NewRequest(stdRequest *http.Request) *Request {
	return &Request{
		stdRequest: stdRequest,
		attributes: map[string]string{},
	}
}

func (r *Request) GetUriPath() string {
	path := r.stdRequest.URL.Path
	ralog.Debug("rahttp: uri path ", path)

	return path
}

func (r *Request) WithMatchedUriPattern(pattern UriPattern) {
	r.matchedUriPattern = pattern
}

func (r *Request) GetMethod() Method {
	return Method(r.stdRequest.Method)
}

func (r *Request) WithAttributes(attributes map[string]string) {
	r.attributes = attributes
}

func (r *Request) GetContentType() string {
	return r.stdRequest.Header.Get("Content-Type")
}

func (r *Request) GetMediaType() string {
	ct := r.GetContentType()
	ralog.Debug("rahttp: Content-Type: ", ct)
	if ct == "" {
		return ""
	}
	s := regexp.MustCompile("\\s*[;,]\\s*").Split(ct, 2)
	mediaType := strings.ToLower(s[0])
	ralog.Debug("rahttp: mediaType: ", mediaType)

	return mediaType
}

func (r *Request) GetParsedBody() map[string]interface{} {
	if r.bodyParsed != nil {
		return r.bodyParsed
	}

	r.bodyParsed = make(map[string]interface{})

	if r.stdRequest.Body == nil {
		return r.bodyParsed
	}

	mt := r.GetMediaType()
	switch {
	case mt == "application/json":
		var reader io.Reader = r.stdRequest.Body
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			ralog.Error("rahttp: errors on reading request body: ", err)
			return r.bodyParsed
		}
		var v interface{}
		ralog.Debug("rahttp: body: ", string(b))
		err = json.Unmarshal(b, &v)
		if err != nil {
			ralog.Informational("rahttp: errors on unmarshalling body: ", err)
			raerror.PanicWith(raerror.TypBadRequest, 0, "Body should be a JSON object")
		}

		r.bodyParsed = formatJsonValue(v)

	default:
		if err := r.stdRequest.ParseForm(); err != nil {
			ralog.Informational("rahttp: errors on parsing form: ", err)
			raerror.PanicWith(raerror.TypBadRequest, 0, "Invalid body format")
		}

		r.bodyParsed = formatUrlValues(r.stdRequest.PostForm)
	}

	return r.bodyParsed
}

func formatUrlValues(postForm url.Values) map[string]interface{} {
	formatedForm := make(map[string]interface{})
	if len(postForm) == 0 {
		return formatedForm
	}

	for key, value := range postForm {
		if len(value) == 1 {
			formatedForm[key] = value[0]
		} else {
			formatedForm[key] = value
		}
	}

	return formatedForm
}

func formatJsonValue(v interface{}) map[string]interface{} {
	m, ok := v.(map[string]interface{})
	ralog.Debug("rahttp: formatted json value: ", m, " assert result: ", ok)
	if !ok {
		ralog.Informational("rahttp: failed to format json value ", v)
		return map[string]interface{}{}
	}

	return m
}

func (r *Request) GetQueryParams() map[string]interface{} {
	if r.queryParams != nil {
		return r.queryParams
	}
	r.queryParams = make(map[string]interface{})

	ralog.Debug("rahttp: URL: ", r.stdRequest.URL)

	if r.stdRequest.URL == nil {
		return r.queryParams
	}

	values := r.stdRequest.URL.Query()
	r.queryParams = formatUrlValues(values)

	ralog.Debug("rahttp: query params: ", r.queryParams)

	return r.queryParams
}

func (r *Request) GetAttributes() map[string]string {
	return r.attributes
}
