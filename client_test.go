package gohttpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx context.Context
}

type TestMethod struct {
	Name, BaseUrl string
	Method        func(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *TestSuite) TearDownSuite() {}

func (s *TestSuite) SetupTest() {}

func (s *TestSuite) TearDownTest() {}

func (s *TestSuite) Test_New_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"

	// Act
	client := New(baseUrl)

	// Assert
	s.NotNil(client)
}

func (s *TestSuite) Test_Request_WhenRequestIsInvalid_ShouldReturnError() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)
	requests := []TestMethod{
		{
			Name:    "GET",
			BaseUrl: baseUrl,
			Method:  client.Get,
		},
		{
			Name:    "POST",
			BaseUrl: baseUrl,
			Method:  client.Post,
		},
		{
			Name:    "PUT",
			BaseUrl: baseUrl,
			Method:  client.Put,
		},
		{
			Name:    "PATCH",
			BaseUrl: baseUrl,
			Method:  client.Patch,
		},
		{
			Name:    "DELETE",
			BaseUrl: baseUrl,
			Method:  client.Delete,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.Name, func() {
			// Act
			response, err := req.Method(nil, "")

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}

func (s *TestSuite) Test_Request_WhenDoReturnsAnError_ShouldReturnError() {
	// Arrange
	baseUrlWithInvalidSchema := "htt \\`"
	client := New(baseUrlWithInvalidSchema)
	requests := []TestMethod{
		{
			Name:    "GET",
			BaseUrl: baseUrlWithInvalidSchema,
			Method:  client.Get,
		},
		{
			Name:    "POST",
			BaseUrl: baseUrlWithInvalidSchema,
			Method:  client.Post,
		},
		{
			Name:    "PUT",
			BaseUrl: baseUrlWithInvalidSchema,
			Method:  client.Put,
		},
		{
			Name:    "PATCH",
			BaseUrl: baseUrlWithInvalidSchema,
			Method:  client.Patch,
		},
		{
			Name:    "DELETE",
			BaseUrl: baseUrlWithInvalidSchema,
			Method:  client.Delete,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.Name, func() {
			// Act
			response, err := req.Method(s.ctx, "")

			// Assert
			s.Nil(response)
			s.Error(err)
		})
	}
}

func (s *TestSuite) Test_Request_ShouldRunSuccesfully() {
	// Arrange
	// init test server
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()

	client := New(svc.URL)
	requests := []TestMethod{
		{
			Name:    "GET",
			BaseUrl: svc.URL,
			Method:  client.Get,
		},
		{
			Name:    "POST",
			BaseUrl: svc.URL,
			Method:  client.Post,
		},
		{
			Name:    "PUT",
			BaseUrl: svc.URL,
			Method:  client.Put,
		},
		{
			Name:    "PATCH",
			BaseUrl: svc.URL,
			Method:  client.Patch,
		},
		{
			Name:    "DELETE",
			BaseUrl: svc.URL,
			Method:  client.Delete,
		},
	}

	for _, req := range requests {
		s.Suite.Run(req.Name, func() {
			// Act
			response, err := req.Method(s.ctx, "")

			// Assert
			s.NotNil(response)
			s.NoError(err)
		})
	}
}
