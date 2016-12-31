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
	if ct == "" {
		return ""
	}
	s := regexp.MustCompile("\\s*[;,]\\s*").Split(ct, 2)
	mediaType := strings.ToLower(s[0])

	return mediaType
}

func (r *Request) GetHeader(key string) string {
	return r.stdRequest.Header.Get(key)
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
			return r.bodyParsed
		}
		var v interface{}
		err = json.Unmarshal(b, &v)
		if err != nil {
			raerror.PanicWith(raerror.TypBadRequest, 0, "rahttp: body should be a JSON object")
		}

		r.bodyParsed = formatJsonValue(v)

	default:
		if err := r.stdRequest.ParseForm(); err != nil {
			raerror.PanicWith(raerror.TypBadRequest, 0, "rahttp: invalid body format")
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
	if !ok {
		return map[string]interface{}{}
	}

	return m
}

func (r *Request) GetQueryParams() map[string]interface{} {
	if r.queryParams != nil {
		return r.queryParams
	}
	r.queryParams = make(map[string]interface{})

	if r.stdRequest.URL == nil {
		return r.queryParams
	}

	values := r.stdRequest.URL.Query()
	r.queryParams = formatUrlValues(values)

	return r.queryParams
}

func (r *Request) GetAttributes() map[string]string {
	return r.attributes
}
