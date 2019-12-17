package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	BaseUrl string
}

type ServiceResponse struct {
	StatusCode int
	Message string
	Data interface {}
}

func (h HttpClient) Get(endpoint string) ServiceResponse {
	response, responseErr := http.Get(fmt.Sprintf("%s%s", h.BaseUrl, endpoint, ))
	defer response.Body.Close()

	return parseResponse(*response, responseErr)
}

func parseResponse(response http.Response, responseErr error) ServiceResponse {
	if responseErr != nil {
		return errorResponse(responseErr.Error())
	}

	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return errorResponse(bodyErr.Error())
	}

	var responseModel interface {}
	unmarshalErr := json.Unmarshal([]byte(body), &responseModel)
	if unmarshalErr != nil {
		return errorResponse(unmarshalErr.Error())
	}

	return ServiceResponse{
		StatusCode: response.StatusCode,
		Message:    "Success",
		Data:       responseModel,
	}
}

func errorResponse(message string) ServiceResponse {
	return ServiceResponse { StatusCode: 400, Message: message, Data: nil}
}