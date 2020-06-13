package go_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

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
	Delete(endpoint string) (*http.Request, error)
	DeleteWith(endpoint string, params interface{}) (*http.Request, error)
	Do(request *http.Request) ServiceResponse
}

// ServiceResponse is a struct who has IsSuccess, StatusCode, Message and Data properties
type ServiceResponse struct {
	IsSuccess  bool
	StatusCode int
	Message    string
	Data       string
}

// New func returns a IHttpClient interface
func New(baseUrl string) IHttpClient {
	return &httpClient{BaseUrl: baseUrl}
}

// Get func returns a request
func (h httpClient) Get(endpoint string) (*http.Request, error) {
	json, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.BaseUrl, endpoint), bytes.NewBuffer(json))
}

// GetWith func returns a request
func (h httpClient) GetWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}

	queryString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.BaseUrl, endpoint, queryString.Encode()), bytes.NewBuffer(json))
}

// Post func returns a request
func (h httpClient) Post(endpoint string) (*http.Request, error) {
	json, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// PostWith func returns a request
func (h httpClient) PostWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// Put func returns a request
func (h httpClient) Put(endpoint string) (*http.Request, error) {
	json, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// PutWith func returns a request
func (h httpClient) PutWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// Delete func returns a request
func (h httpClient) Delete(endpoint string) (*http.Request, error) {
	json, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// DeleteWith func returns a request
func (h httpClient) DeleteWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// Do func returns a response with your data
func (h httpClient) Do(request *http.Request) ServiceResponse {
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ServiceResponse{
			StatusCode: -1,
			Message:    "An error occured while doing the request!",
		}
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ServiceResponse{
			StatusCode: -1,
			Message:    "An error occured while reading the body of response!",
		}
	}

	return ServiceResponse{
		IsSuccess:  true,
		StatusCode: response.StatusCode,
		Message:    "Success",
		Data:       string(body),
	}
}