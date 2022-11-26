package main

import (
	"context"
	"log"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	ctx := context.Background()

	opts := []gohttpclient.ClientOption{
		gohttpclient.WithDefaultHeaders(),
		gohttpclient.WithTimeout(time.Second * 3),
	}
	client := gohttpclient.New("https://jsonplaceholder.typicode.com", opts...)

	reqOpts := []gohttpclient.Option{
		gohttpclient.WithHeader("x-useragent", "go-http-client"),
		gohttpclient.WithHeader("x-correlationid", "123456789"),
	}
	response, err := client.Get(ctx, "/posts/1", reqOpts...)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var post Post
	if err := response.Unmarshal(&post); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf(post.Title) // sunt aut facere repellat provident occaecati...
}
