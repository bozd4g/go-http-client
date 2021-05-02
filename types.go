package gohttpclient

import "net/http"

// client is a struct who has BaseUrl property
type client struct {
	BaseUrl string
}

// Client is a interface who calls the methods
type Client interface {
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
	Do(request *http.Request) (Response, error)
}

// ResponseStruct is a struct who returns after requests
type ResponseStruct struct {
	Status        string
	StatusCode    int
	Header        http.Header
	ContentLength int64
	Body          []byte
}

// Response is an interface of ResponseStruct struct
type Response interface {
	Get() ResponseStruct
	To(value interface{})
}
