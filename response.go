package gohttpclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	Response struct {
		res  *http.Response
		body []byte
	}
)

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) Unmarshal(v any) error {
	defer r.res.Body.Close()
	body, err := ioutil.ReadAll(r.res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func (r *Response) Status() int {
	return r.res.StatusCode
}

func (r *Response) Headers() http.Header {
	return r.res.Header
}

func (r *Response) Cookies() []*http.Cookie {
	return r.res.Cookies()
}

func (r *Response) Ok() bool {
	return r.res.StatusCode >= 200 && r.res.StatusCode <= 299
}

func (r *Response) Get() *http.Response {
	return r.res
}
