package gohttpclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestClientSuite struct {
	suite.Suite
	ctx context.Context
}

type TestMethod struct {
	name, baseUrl string
	method        func(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
	options       []Option
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}

func (s *TestClientSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *TestClientSuite) Test_New_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"

	// Act
	client := New(baseUrl)

	// Assert
	s.NotNil(client)
}

func (s *TestClientSuite) Test_Request_WhenRequestIsInvalid_ShouldReturnError() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)
	requests := []TestMethod{
		{
			name:    "GET",
			baseUrl: baseUrl,
			method:  client.Get,
		},
		{
			name:    "POST",
			baseUrl: baseUrl,
			method:  client.Post,
		},
		{
			name:    "PUT",
			baseUrl: baseUrl,
			method:  client.Put,
		},
		{
			name:    "PATCH",
			baseUrl: baseUrl,
			method:  client.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: baseUrl,
			method:  client.Delete,
		},
		{
			name:    "CONNECT",
			baseUrl: baseUrl,
			method:  client.Connect,
		},
		{
			name:    "OPTIONS",
			baseUrl: baseUrl,
			method:  client.Options,
		},
		{
			name:    "TRACE",
			baseUrl: baseUrl,
			method:  client.Trace,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.name, func() {
			// Act
			response, err := req.method(nil, "")

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}

func (s *TestClientSuite) Test_Request_WhenDoReturnsAnError_ShouldReturnError() {
	// Arrange
	baseUrlWithInvalidSchema := "htt \\`"
	client := New(baseUrlWithInvalidSchema)
	requests := []TestMethod{
		{
			name:    "GET",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Get,
		},
		{
			name:    "POST",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Post,
		},
		{
			name:    "PUT",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Put,
		},
		{
			name:    "PATCH",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Delete,
		},
		{
			name:    "CONNECT",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Connect,
		},
		{
			name:    "OPTIONS",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Options,
		},
		{
			name:    "TRACE",
			baseUrl: baseUrlWithInvalidSchema,
			method:  client.Trace,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.name, func() {
			// Act
			response, err := req.method(s.ctx, "")

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}

func (s *TestClientSuite) Test_Request_WhenBodyReturnsError_ShouldReturnError() {
	// Arrange
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
		w.Write(nil)
	}))
	defer svc.Close()

	client := New(svc.URL)
	requests := []TestMethod{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  client.Get,
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  client.Post,
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  client.Put,
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  client.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  client.Delete,
		},
		{
			name:    "CONNECT",
			baseUrl: svc.URL,
			method:  client.Connect,
		},
		{
			name:    "OPTIONS",
			baseUrl: svc.URL,
			method:  client.Options,
		},
		{
			name:    "TRACE",
			baseUrl: svc.URL,
			method:  client.Trace,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.name, func() {
			// Act
			response, err := req.method(s.ctx, "")

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}

func (s *TestClientSuite) Test_Request_ShouldRunSuccesfully() {
	// Arrange
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()

	client := New(svc.URL)
	requests := []TestMethod{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  client.Get,
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  client.Post,
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  client.Put,
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  client.Patch,
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  client.Delete,
		},
		{
			name:    "CONNECT",
			baseUrl: svc.URL,
			method:  client.Connect,
		},
		{
			name:    "OPTIONS",
			baseUrl: svc.URL,
			method:  client.Options,
		},
		{
			name:    "TRACE",
			baseUrl: svc.URL,
			method:  client.Trace,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.name, func() {
			// Act
			response, err := req.method(s.ctx, "")

			// Assert
			s.NotNil(response)
			s.NoError(err)
		})
	}
}

func (s *TestClientSuite) Test_Request_WithOptions_ShouldRunSuccesfully() {
	// Arrange
	// init test server
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()

	client := New(svc.URL)
	requests := []TestMethod{
		{
			name:    "GET",
			baseUrl: svc.URL,
			method:  client.Get,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "POST",
			baseUrl: svc.URL,
			method:  client.Post,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "PUT",
			baseUrl: svc.URL,
			method:  client.Put,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "PATCH",
			baseUrl: svc.URL,
			method:  client.Patch,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "DELETE",
			baseUrl: svc.URL,
			method:  client.Delete,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "CONNECT",
			baseUrl: svc.URL,
			method:  client.Connect,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "OPTIONS",
			baseUrl: svc.URL,
			method:  client.Options,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
		{
			name:    "TRACE",
			baseUrl: svc.URL,
			method:  client.Trace,
			options: []Option{WithHeader("key", "value"), WithQuery("key", "value")},
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.name, func() {
			// Act
			response, err := req.method(s.ctx, "", req.options...)

			// Assert
			s.NotNil(response)
			s.NoError(err)
		})
	}
}

func (s *TestClientSuite) Test_PrepareRequest_WhenRequestIsInvalid_ShouldReturnError() {
	// Arrange
	baseUrl := "http://localhost:8080"
	method := "GET"
	endpoint := "/test"
	client := New(baseUrl)

	// Act
	response, err := client.PrepareRequest(nil, method, endpoint)

	// Assert
	s.Nil(response)
	s.Error(err)
}

func (s *TestClientSuite) Test_PrepareRequest_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"
	method := "GET"
	endpoint := "/test"
	body := []byte("test")
	client := New(baseUrl)

	// Act
	request, err := client.PrepareRequest(s.ctx, method, endpoint, WithBody(body))

	// Assert
	s.NoError(err)
	s.NotNil(request)
	s.Equal(method, request.Method)
	s.Equal(baseUrl+endpoint, request.URL.String())

	requestBody, err := ioutil.ReadAll(request.Body)
	s.NoError(err)
	s.Equal(body, requestBody)
}
