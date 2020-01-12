
<p align="center">
  <a href="https://github.com/bozd4g/go-http-client">
    <img alt="go-http-client" src="https://raw.githubusercontent.com/bozd4g/go-http-client/master/assets/github/logo.png" width="500">
  </a>
</p>

<h1 align="center">
  go-http-client
</h1>

<p align="center">
  An enhanced http client for <a href="https://golang.org/">Golang</a>
</p>

<p align="center">
  <a href="https://godoc.org/github.com/bozd4g/go-http-client/client" target="_blank">Documentation on GoDoc 🔗</a>
</p>

<p align="center">
  <a href="https://bozd4g.mit-license.org/"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
  <a href="https://travis-ci.org/bozd4g/go-http-client"><img src="https://travis-ci.org/bozd4g/go-http-client.svg?branch=master"></a>
</p>

This package provides you a http client package for your http requests. You can send requests quicly with this package. If you want to contribute this package, please fork and [create](https://github.com/bozd4g/go-http-client/pulls) a pull request.

# Installation

```
$ go get github.com/bozd4g/go-http-client/
```

# Usage
```go
import (
	"fmt"
	"github.com/bozd4g/go-http-client/client"
	"github.com/mitchellh/mapstructure"
)

type Todo struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

func main() {
	client := client.HttpClient{BaseUrl: "http://jsonplaceholder.typicode.com"}
	response := client.PostWithParameters("/posts", Todo{
		Id:        1,
		UserId:    1,
		Title:     "Lorem ipsum dolor sit amet",
		Completed: true,
	})

	if response.IsSuccess {
		var todo Todo
		mapstructure.Decode(response.Data, &todo)
		fmt.Println(todo.Title) // Lorem ipsum dolor sit amet

	} else {
		fmt.Println(response.Message)
	}
}

```

## Functions

All functions return a type called ServiceResponse as below.
```go
type ServiceResponse struct {
	IsSuccess bool
	StatusCode int
	StatusText string
	Message string
	Data interface {}
}
```

You can call these functions from your application.

| Function                                                  | Has Params |
| --------------------------------------------------------- | ---------- |
| Get(endpoint string)                                      | -          |
| GetWithParameters(endpoint string, params interface {})   | Yes        |
| Post(endpoint string)                                     | -          |
| PostWithParameters(endpoint string, params interface {})  | Yes        |
| Put(endpoint string)                                      | -          |
| PutWithParameters(endpoint string, params interface{})    | Yes        |
| Delete(endpoint string)                                   | -          |
| DeleteWithParameters(endpoint string, params interface{}) | Yes        |

# License
Copyright (c) 2020 Furkan Bozdag

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
