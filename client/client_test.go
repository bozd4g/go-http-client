package client

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

func TestGetRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://go-http-client.free.beeceptor.com" }
	response := client.Get("/posts")

	t.Run("Returns a todo who have id as 1", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 101

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPostRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://go-http-client.free.beeceptor.com" }
	response := client.PostWithParameters("/posts", Todo{
		Id: 1,
	})

	t.Run("Returns a todo who have id as 101", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 201

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestPutRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://go-http-client.free.beeceptor.com" }
	response := client.PutWithParameters("/posts", Todo{
		Id: 1,
	})

	t.Run("Returns a todo who have id as 101", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 301

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}

func TestDeleteRequest (t *testing.T) {
	client := client.HttpClient { BaseUrl: "https://go-http-client.free.beeceptor.com" }
	response := client.DeleteWithParameters("/posts", Todo{
		Id: 1,
	})

	t.Run("Returns a todo who have id as 101", func(t *testing.T) {
		var got Todo
		mapstructure.Decode(response.Data, &got)
		want := 401

		if got.Id != want {
			t.Errorf("Unexpected data. Got: %d, expected: %d", got.Id, want)
		}
	})
}