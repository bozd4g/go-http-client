package gohttpclient

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestOptionSuite struct {
	suite.Suite
	ctx context.Context
}

func TestOption(t *testing.T) {
	suite.Run(t, new(TestOptionSuite))
}

func (s *TestOptionSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *TestOptionSuite) Test_WithDefaultHeaders_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)

	// Act
	WithDefaultHeaders()(client)

	// Assert
	s.Assert().Equal("application/json", client.headers["Content-Type"].Value)
	s.Assert().Equal("application/json", client.headers["Accept"].Value)
}

func (s *TestOptionSuite) Test_WithDefaultHeaders_WhenCalledByNew_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"

	// Act
	client := New(baseUrl, WithDefaultHeaders())

	// Assert
	s.Assert().Equal("application/json", client.headers["Content-Type"].Value)
	s.Assert().Equal("application/json", client.headers["Accept"].Value)
}

func (s *TestOptionSuite) Test_WithTimeout_ShouldRunSuccesfully() {
	// Arrange
	timeout := 5 * time.Second
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)

	// Act
	WithTimeout(timeout)(client)

	// Assert
	s.Assert().Equal(timeout, client.timeout)
}

func (s *TestOptionSuite) Test_WithHeader_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)

	// Act
	WithHeader("Content-Type", "application/json")(client)

	// Assert
	s.Assert().Equal("application/json", client.headers["Content-Type"].Value)
}

func (s *TestOptionSuite) Test_WithQuery_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)

	// Act
	WithQuery("key", "value")(client)

	// Assert
	s.Assert().Equal("value", client.query["key"])
}

func (s *TestOptionSuite) Test_WithBody_ShouldRunSuccesfully() {
	// Arrange
	baseUrl := "http://localhost:8080"
	client := New(baseUrl)

	// Act
	WithBody([]byte("body"))(client)

	// Assert
	s.Assert().Equal("body", string(client.body))
}
