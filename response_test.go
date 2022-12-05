package gohttpclient

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockReadCloser struct {
	mock.Mock
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

type TestResponseSuite struct {
	suite.Suite
	ctx context.Context
}

func TestResponse(t *testing.T) {
	suite.Run(t, new(TestResponseSuite))
}

func (s *TestResponseSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *TestResponseSuite) Test_Body_ShouldRunSuccesfully() {
	// Arrange
	body := []byte("test")
	res := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	// Act
	resp := Response{res, body}

	// Assert
	s.Equal(body, resp.Body())
}

func (s *TestResponseSuite) Test_Unmarshal_ShouldRunSuccesfully() {
	// Arrange
	body := []byte(`{"name":"test"}`)

	// Act
	resp := Response{
		res: &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer(body)),
		},
	}

	// Assert
	var response map[string]interface{}
	err := resp.Unmarshal(&response)
	s.NoError(err)
	s.Equal("test", response["name"])
}

func (s *TestResponseSuite) Test_Unmarshal_WhenBodyIsWrong_ShouldReturnError() {
	// Arrange
	mockReadCloser := mockReadCloser{}
	mockReadCloser.On("Read", mock.AnythingOfType("[]uint8")).Return(0, fmt.Errorf("error reading"))
	mockReadCloser.On("Close").Return(fmt.Errorf("error closing"))

	resp := Response{
		res: &http.Response{
			Body: &mockReadCloser,
		},
	}

	// Act
	var response map[string]interface{}
	err := resp.Unmarshal(&response)

	// Assert
	s.Errorf(err, "error reading")
}

func (s *TestResponseSuite) Test_Unmarshal_WhenUnMarshalReturnsError_ShouldReturnError() {
	// Arrange
	body := []byte(`{"name":"test"`)

	// Act
	resp := Response{
		res: &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer(body)),
		},
	}

	// Assert
	var response map[string]interface{}
	err := resp.Unmarshal(&response)
	s.Error(err)
}

func (s *TestResponseSuite) Test_Status_ShouldRunSuccesfully() {
	// Arrange
	resp := Response{
		res: &http.Response{
			StatusCode: 200,
		},
	}

	// Act
	status := resp.Status()

	// Assert
	s.Equal(200, status)
}

func (s *TestResponseSuite) Test_Header_ShouldRunSuccesfully() {
	// Arrange
	resp := Response{
		res: &http.Response{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}

	// Act
	header := resp.Headers()

	// Assert
	s.Equal("application/json", header["Content-Type"][0])
}

func (s *TestResponseSuite) Test_Cookies_ShouldRunSuccesfully() {
	// Arrange
	resp := Response{
		res: &http.Response{
			Header: http.Header{
				"Set-Cookie": []string{"test=1"},
			},
		},
	}

	// Act
	cookies := resp.Cookies()

	// Assert
	s.Equal("test=1", cookies[0].String())
}

func (s *TestResponseSuite) Test_Ok_ShouldRunSuccesfully() {
	// Arrange
	resp := Response{
		res: &http.Response{
			StatusCode: 200,
		},
	}

	// Act
	ok := resp.Ok()

	// Assert
	s.True(ok)
}

func (s *TestResponseSuite) Test_Get_ShouldRunSuccesfully() {
	// Arrange
	resp := Response{
		res: &http.Response{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}

	// Act
	res := resp.Get()

	// Assert
	s.Equal(resp.res, res)
}
