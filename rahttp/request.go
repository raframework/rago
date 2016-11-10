package rahttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/raframework/rago/ralog"
)

type Request struct {
	stdRequest        *http.Request
	matchedUriPattern UriPattern
	attributes        map[string]string
}

func NewRequest(stdRequest *http.Request) *Request {
	return &Request{
		stdRequest: stdRequest,
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
	if ct == "" {
		return ""
	}
	s := regexp.MustCompile("/\\s*[;,]\\s*/").Split(ct, 2)

	return s[0]
}

func (r *Request) GetParsedBody() map[string]interface{} {
	parsedBody := make(map[string]interface{})

	if r.stdRequest.Body == nil {
		return parsedBody
	}

	mt := r.GetMediaType()
	switch {
	case mt == "application/json":
		var reader io.Reader = r.stdRequest.Body
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		var v interface{}
		ralog.Debug("rahttp: body: ", string(b))
		err = json.Unmarshal(b, &v)
		if err != nil {
			panic("Body should be a JSON object: " + err.Error())
		}
		parsedBody = formatJsonValue(v)

	default:
		r.stdRequest.ParseForm()
		parsedBody = formatPostForm(r.stdRequest.PostForm)
	}

	return parsedBody
}

func formatPostForm(postForm url.Values) map[string]interface{} {
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
	log.Println("rahttp: formatted json value: ", m, " assert result: ", ok)
	if !ok {
		return map[string]interface{}{}
	}

	return m
}

func (r *Request) GetQueryParams() map[string]string {
	return nil
}

func (r *Request) GetAttributes() map[string]string {
	return nil
}
