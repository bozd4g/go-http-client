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
		return ServiceResponse { StatusCode:500, Message: responseErr.Error(), Data: nil}
	}

	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return ServiceResponse { StatusCode:500, Message: bodyErr.Error(), Data: nil}
	}

	var responseModel interface {}
	error := json.Unmarshal([]byte(body), &responseModel)
	if error != nil {
		return ServiceResponse { StatusCode:500, Message: error.Error(), Data: nil}
	}

	return ServiceResponse{
		StatusCode: response.StatusCode,
		Message:    "Success",
		Data:       responseModel,
	}
}