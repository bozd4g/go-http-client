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

func TestGetRequestWith(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.GetWith("posts", Todo{
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

func TestGetRequestWithAsWrong(t *testing.T) {
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

func TestPostRequestWith(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.PostWith("posts", Todo{
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

func TestPostRequestWithAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.PostWith("qeqwe", Todo{
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

func TestPutRequestWith(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.PutWith("posts", Todo{
		Id: 31,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestPutRequestWithAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.PutWith("", Todo{
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

func TestDeleteRequestWith(t *testing.T) {
	client := client.HttpClient{BaseUrl: "https://jsonplaceholder.typicode.com/"}
	response := client.DeleteWith("posts", Todo{
		Id: 41,
	})

	t.Run("Returns 200 response", func(t *testing.T) {
		if response.IsSuccess == false {
			t.Errorf("Unexpected data. Got: %t, expected: %t", false, true)
		}
	})
}

func TestDeleteRequestWithAsWrong(t *testing.T) {
	client := client.HttpClient{BaseUrl: ""}
	response := client.DeleteWith("", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if response.StatusCode != 2 {
			t.Errorf("Unexpected data. Got: %v, expected: %v", response.StatusCode, 2)
		}
	})
}
