package go_http_client

import "net/http"

// httpClient is a struct who has BaseUrl property
type httpClient struct {
	BaseUrl string
}

// IHttpClient is a interface who calls the methods
type IHttpClient interface {
	Get(endpoint string) (*http.Request, error)
	GetWith(endpoint string, params interface{}) (*http.Request, error)
	Post(endpoint string) (*http.Request, error)
	PostWith(endpoint string, params interface{}) (*http.Request, error)
	Put(endpoint string) (*http.Request, error)
	PutWith(endpoint string, params interface{}) (*http.Request, error)
	Patch(endpoint string) (*http.Request, error)
	PatchWith(endpoint string, params interface{}) (*http.Request, error)
	Delete(endpoint string) (*http.Request, error)
	DeleteWith(endpoint string, params interface{}) (*http.Request, error)
	Do(request *http.Request) (IHttpResponse, error)
}

// HttpResponse is a struct who returns after requests
type HttpResponse struct {
	Status        string
	StatusCode    int
	Header        http.Header
	ContentLength int64
	Body          []byte
}

// IHttpResponse is an interface of HttpResponse struct
type IHttpResponse interface {
	Get() HttpResponse
	To(value interface{})
}
