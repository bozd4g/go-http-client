package gohttpclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	Response struct {
		httpResponse *http.Response
	}
)

func (r *Response) Body() ([]byte, error) {
	defer r.httpResponse.Body.Close()
	return ioutil.ReadAll(r.httpResponse.Body)
}

func (r *Response) Json(v any) error {
	defer r.httpResponse.Body.Close()
	body, err := ioutil.ReadAll(r.httpResponse.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func (r *Response) Status() int {
	return r.httpResponse.StatusCode
}

func (r *Response) Header() http.Header {
	return r.httpResponse.Header
}

func (r *Response) Cookies() []*http.Cookie {
	return r.httpResponse.Cookies()
}

func (r *Response) Ok() bool {
	return r.httpResponse.StatusCode >= 200 && r.httpResponse.StatusCode <= 299
}
