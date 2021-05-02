
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
  <a href="https://pkg.go.dev/github.com/bozd4g/go-http-client" target="_blank">Documentation on go.dev 🔗</a>
</p>

<p align="center">
  <a href="https://bozd4g.mit-license.org/"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
  <a href="https://travis-ci.org/bozd4g/go-http-client"><img src="https://travis-ci.org/bozd4g/go-http-client.svg?branch=master"></a>
  <a href="https://goreportcard.com/report/github.com/bozd4g/go-http-client"><img src="https://goreportcard.com/badge/github.com/bozd4g/go-http-client"></a>
<a href="https://codecov.io/gh/bozd4g/go-http-client">
<img alt="Coverage" src="https://codecov.io/gh/bozd4g/go-http-client/branch/master/graphs/badge.svg?branch=master">
</a>
<a href="https://github.com/avelino/awesome-go">
<img alt="awesome" src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg">
</a>
</p>

This package provides you a http client package for your http requests. You can send requests quicly with this package. If you want to contribute this package, please fork and [create](https://github.com/bozd4g/go-http-client/pulls) a pull request.

# Installation
```
$ go get -u github.com/bozd4g/go-http-client/
```

# Usage
```go
import (
	"encoding/json"
	"fmt"
	client "github.com/bozd4g/go-http-client"
)

type Todo struct {
    Id        int
    UserId    int
    Title     string
    Completed bool
}

func main() {
    httpClient := client.New("https://jsonplaceholder.typicode.com/")
    request, err := httpClient.Get("posts/10")
    
    if err != nil {
        panic(err)
    }
    
    response, err := httpClient.Do(request)
    if err != nil {
        panic(err)
    }
    
    var todo Todo
    err = json.Unmarshal(response.Get().Body, &todo)
    if err != nil {
        panic(err)
    }
    fmt.Println(todo.Title) // Lorem ipsum dolor sit amet

    // or  
    var todo2 Todo     
    response, err = httpClient.Do(request)
    if err == nil {
        response.To(&todo2)
        fmt.Println(todo2.Title) // Lorem ipsum dolor sit amet
    } else {
        fmt.Println(err.Error())
    }
}

```

## Functions
You can call these functions from your application.

| Function                                                  | Has Params |
| --------------------------------------------------------- | ---------- |
| Get(endpoint string)                                      | - |
| GetWith(endpoint string, params interface {})   | Yes        |
| Post(endpoint string)                                     | - |
| PostWith(endpoint string, params interface {})  | Yes        |
| Patch(endpoint string)                                      | - |
| PatchWith(endpoint string, params interface{})    | Yes        |
| Put(endpoint string)                                      | - |
| PutWith(endpoint string, params interface{})    | Yes        |
| Delete(endpoint string)                                   | - |
| DeleteWith(endpoint string, params interface{}) | Yes         |
| Do() (Response, error) | - |
| To(value interface{}) | - |

# License
Copyright (c) 2020 Furkan Bozdag

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
