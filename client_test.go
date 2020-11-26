package go_http_client_test

import (
	"encoding/json"
	client "github.com/bozd4g/go-http-client"
	"testing"
)

type Todo struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

func TestGetRequest(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Get("posts/10")

	t.Run("Returns a todo who have id as 10", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 10
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestGetRequestAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Get("posts/123123")

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestGetRequestWith(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.GetWith("posts", Todo{
		Id: 11,
	})

	t.Run("Returns a todo who have id as 11", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got []Todo
		want := 11
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if len(got) < 1 || got[10].Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got[10].Id, want)
		}
	})
}

func TestGetRequestWithAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.GetWith("posts/123123", nil)

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestGetRequestWithTo(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Get("posts/10")

	t.Run("Returns a todo who have id as 10", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		var got Todo
		want := 10
		_, err := httpClient.Do(request).To(&got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPostRequest(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Post("posts/20")

	t.Run("Returns a todo who have id as 20", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 0
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPostRequestAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Post("posts/23213123")

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPostRequestWith(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PostWith("posts", Todo{
		Id: 21,
	})

	t.Run("Returns a todo who have id as 21", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 101
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPostRequestWithAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PostWith("qeqwe", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPutRequest(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Put("posts/30")

	t.Run("Returns a todo who have id as 30", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 30
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequestAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Put("posts/23213123")

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPutRequestWith(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PutWith("posts", Todo{
		Id: 31,
	})

	t.Run("Returns a todo who have id as 31", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 0
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequestWithAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PostWith("qeqwe", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPatchRequest(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Patch("posts/22")

	t.Run("Returns a todo who have id as 22", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 22
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPatchRequestAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Patch("posts/45156416546")

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestPatchRequestWith(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PatchWith("posts", Todo{
		Id: 33,
	})

	t.Run("Returns a todo who have id as 33", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 0
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPatchRequestWithAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.PatchWith("qeqwe", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestDeleteRequest(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Delete("posts/40")

	t.Run("Returns a todo who have id as 40", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 0
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestDeleteRequestAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.Delete("posts/23213123")

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}

func TestDeleteRequestWith(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.DeleteWith("posts", Todo{
		Id: 41,
	})

	t.Run("Returns a todo who have id as 41", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}

		var got Todo
		want := 0
		err := json.Unmarshal([]byte(response.Data), &got)
		if err != nil {
			t.Error(err.Error())
		}

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestDeleteRequestWithAsWrong(t *testing.T) {
	httpClient := client.New("https://jsonplaceholder.typicode.com/")
	request, err := httpClient.DeleteWith("qeqwe", Todo{
		Id: -111111111,
	})

	t.Run("Returns an error", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		response := httpClient.Do(request)
		if !response.IsSuccess {
			t.Errorf("Unexpected data. Got: %v, expected: %v", true, false)
		}
	})
}
