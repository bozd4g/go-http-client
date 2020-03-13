package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

// HttpClient is a struct who has BaseUrl property
type HttpClient struct {
	BaseUrl string
}

// ServiceResponse is a struct who has IsSuccess, StatusCode and their text, Message and Data properties
type ServiceResponse struct {
	IsSuccess  bool
	StatusCode ResponseType
	StatusText string
	Message    string
	Data       interface{}
}

// ResponseType is a enum who has Success, InternalError and ServerError properties
type ResponseType int
const (
	// Success says the request is completed successfully
	Success       ResponseType = 0
	// InternalError says the request is completed with internal errors
	InternalError ResponseType = 1
	// ServerError says the request is completed with server errors
	ServerError   ResponseType = 2
)

// Text func returns a string of ResponseType
func (r ResponseType) Text() string {
	if r == InternalError {
		return "Request Error"
	}

	return "Server Error"
}

// Get func returns a response with your data
func (h HttpClient) Get(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.BaseUrl, endpoint), bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// GetWithParameters func returns a response with your data
func (h HttpClient) GetWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	queryString, _ := query.Values(params)

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.BaseUrl, endpoint, queryString.Encode()), bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// Post func returns a response with your data
func (h HttpClient) Post(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// PostWithParameters func returns a response with your data
func (h HttpClient) PostWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// Put func returns a response with your data
func (h HttpClient) Put(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// PutWithParameters func returns a response with your data
func (h HttpClient) PutWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// Delete func returns a response with your data
func (h HttpClient) Delete(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// DeleteWithParameters func returns a response with your data
func (h HttpClient) DeleteWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

func parseResponse(request *http.Request, requestErr error) ServiceResponse {
	if requestErr != nil {
		return errorResponse(requestErr.Error(), InternalError)
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, responseErr := client.Do(request)

	if responseErr != nil {
		return errorResponse(responseErr.Error(), ServerError)
	}

	defer response.Body.Close()

	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return errorResponse(bodyErr.Error(), InternalError)
	}

	var responseModel interface{}
	unmarshalErr := json.Unmarshal([]byte(string(body)), &responseModel)
	if unmarshalErr != nil {
		return errorResponse(unmarshalErr.Error(), InternalError)
	}

	return ServiceResponse{
		IsSuccess:  true,
		StatusCode: Success,
		StatusText: http.StatusText(response.StatusCode),
		Data:       responseModel,
		Message:    "Success",
	}
}

func errorResponse(message string, statusCode ResponseType) ServiceResponse {
	return ServiceResponse{
		IsSuccess:  false,
		StatusCode: statusCode,
		StatusText: statusCode.Text(),
		Message:    message,
		Data:       nil}
}
