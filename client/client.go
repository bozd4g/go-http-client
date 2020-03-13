package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

// The class object as below
type HttpClient struct {
	BaseUrl string
}

// The response type of service requests
type ServiceResponse struct {
	IsSuccess  bool
	StatusCode int
	StatusText string
	Message    string
	Data       interface{}
}

// It returns your GET response with your data
func (h HttpClient) Get(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.BaseUrl, endpoint), bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your GET response with your data
func (h HttpClient) GetWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	queryString, _ := query.Values(params)

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.BaseUrl, endpoint, queryString.Encode()), bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your POST response with your data
func (h HttpClient) Post(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your POST response with your data
func (h HttpClient) PostWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your PUT response with your data
func (h HttpClient) Put(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your POST response with your data
func (h HttpClient) PutWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your DELETE response with your data
func (h HttpClient) Delete(endpoint string) ServiceResponse {
	json, _ := json.Marshal(map[string]string{})
	request, requestErr := http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

// It returns your DELETE response with your data
func (h HttpClient) DeleteWithParameters(endpoint string, params interface{}) ServiceResponse {
	json, _ := json.Marshal(params)
	request, requestErr := http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))

	return parseResponse(request, requestErr)
}

func parseResponse(request *http.Request, requestErr error) ServiceResponse {
	if requestErr != nil {
		return errorResponse(requestErr.Error())
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return errorResponse(responseErr.Error())
	}

	defer response.Body.Close()

	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return errorResponse(bodyErr.Error())
	}

	var responseModel interface{}
	unmarshalErr := json.Unmarshal([]byte(string(body)), &responseModel)
	if unmarshalErr != nil {
		return errorResponse(unmarshalErr.Error())
	}

	return ServiceResponse{
		IsSuccess:  true,
		StatusCode: response.StatusCode,
		StatusText: http.StatusText(response.StatusCode),
		Data:       responseModel,
		Message:    "Success",
	}
}

func errorResponse(message string) ServiceResponse {
	return ServiceResponse{
		IsSuccess:  false,
		StatusCode: http.StatusBadRequest,
		StatusText: http.StatusText(http.StatusBadRequest),
		Message:    message,
		Data:       nil}
}
