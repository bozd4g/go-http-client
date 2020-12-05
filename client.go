package go_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

// New func returns a IHttpClient interface
func New(baseUrl string) IHttpClient {
	return &httpClient{BaseUrl: baseUrl}
}

// Get func returns a request
func (h httpClient) Get(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.BaseUrl, endpoint), bytes.NewBuffer([]byte{}))
}

// GetWith func returns a request
func (h httpClient) GetWith(endpoint string, params interface{}) (*http.Request, error) {
	queryString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.BaseUrl, endpoint, queryString.Encode()), bytes.NewBuffer([]byte{}))
}

// Post func returns a request
func (h httpClient) Post(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
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
	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// PutWith func returns a request
func (h httpClient) PutWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// Patch func returns a request
func (h httpClient) Patch(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPatch, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// PatchWith func returns a request
func (h httpClient) PatchWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

// Delete func returns a request
func (h httpClient) Delete(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
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
func (h httpClient) Do(request *http.Request) (IHttpResponse, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Header:        response.Header,
		ContentLength: response.ContentLength,
		Body:          body,
	}, nil
}

// Get func returns HttpResponse struct of request
func (r HttpResponse) Get() HttpResponse {
	return r
}

// To func returns converts string to struct
func (r HttpResponse) To(value interface{}) {
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
	}
}
