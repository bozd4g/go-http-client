package main

import (
	ghc "github.com/bozd4g/go-http-client"
)

func main() {
	_ = ghc.New("https://jsonplaceholder.typicode.com")
}
