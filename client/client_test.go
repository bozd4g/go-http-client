package client_test

import (
	"github.com/bozd4g/go-http-client/client"
	"github.com/mitchellh/mapstructure"
	"testing"
)

type Todo struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

func TestGetRequest(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.Get("posts/10")

	t.Run("Returns a todo who have id as 10", func(t *testing.T) {
		var got Todo
		want := 10
		err := mapstructure.Decode(response.Data, &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestGetRequestAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.Get("posts/123123")

	t.Run("Returns an error", func(t *testing.T) {
		if response.IsSuccess == true {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestGetRequestWithParameters(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.GetWithParameters("posts", Todo{
		Id: 11,
	})

	t.Run("Returns a todo who have id as 11", func(t *testing.T) {
		var got []Todo
		want := 11
		err := mapstructure.Decode(response.Data, &got)
		if err != nil {
			t.Error(err.Error())
		}

		if len(got) < 1 || got[10].Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got[10].Id, want)
		}
	})
}

func TestGetRequestWithParametersAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.Get("posts/123123")

	t.Run("Returns an error", func(t *testing.T) {
		if response.StatusCode != 2 {
			t.Errorf("Unexpected data. Got: %v, expected: %v", response.StatusCode, 2)
		}
	})
}

func TestPostRequest(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.Post("posts/20")

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestPostRequestAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.Post("posts/23213123")

	t.Run("Returns an error", func(t *testing.T) {
		if response.IsSuccess == true {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPostRequestWithParameters(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.PostWithParameters("posts", Todo{
		Id: 21,
	})

	t.Run("Returns a todo who have id as 20", func(t *testing.T) {
		var got Todo
		want := 21
		err := mapstructure.Decode(response.Data, &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPostRequestWithParametersAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.PostWithParameters("qeqwe", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if response.StatusCode != 2 {
			t.Errorf("Unexpected data. Got: %v, expected: %v", response.StatusCode, 2)
		}
	})
}

func TestPutRequest(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.Put("posts/30")

	t.Run("Returns a todo who have id as 30", func(t *testing.T) {
		var got Todo
		want := 30
		err := mapstructure.Decode(response.Data, &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequestAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.Put("posts/asdasd")

	t.Run("Returns an error", func(t *testing.T) {
		if response.IsSuccess == true {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPutRequestWithParameters(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.PutWithParameters("posts", Todo{
		Id: 31,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestPutRequestWithParametersAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.PutWithParameters("", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if response.StatusCode != 2 {
			t.Errorf("Unexpected data. Got: %v, expected: %v", response.StatusCode, 2)
		}
	})
}

func TestDeleteRequest(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.Delete("posts/40")

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestDeleteRequestAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.Delete("posts/asdasd")

	t.Run("Returns an error", func(t *testing.T) {
		if response.IsSuccess == true {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestDeleteRequestWithParameters(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.DeleteWithParameters("posts", Todo{
		Id: 41,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestDeleteRequestWithParametersAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.DeleteWithParameters("", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if response.StatusCode != 2 {
			t.Errorf("Unexpected data. Got: %v, expected: %v", response.StatusCode, 2)
		}
	})
}