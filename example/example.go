package main

import (
	"context"
	"log"

	gohttpclient "github.com/bozd4g/go-http-client"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	ctx := context.Background()
	client := gohttpclient.New("https://jsonplaceholder.typicode.com")

	response, err := client.Get(ctx, "/posts/1")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var post Post
	if err := response.Json(&post); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf(post.Title) // sunt aut facere repellat provident occaecati...
}
