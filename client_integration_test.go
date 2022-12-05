package gohttpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	bigJsonFile []byte
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupSuite() {
	// read large file from test folder

}

func (s *IntegrationSuite) Test_Get_WhenServerReturnsBigFile_ShouldRunSuccesfully() {
	// Arrange
	type Post struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	ctx := context.Background()

	// read large file from test folder
	testSvc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bigArr := make([]Post, 0)
		for i := 0; i < 100000; i++ {
			bigArr = append(bigArr, Post{ID: i, Title: fmt.Sprintf("Title %d", i)})
		}

		postJson, err := json.Marshal(bigArr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write cookie
		http.SetCookie(w, &http.Cookie{Name: "test", Value: "test"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(postJson)
	}))

	client := New(testSvc.URL)

	// Act
	resp, err := client.Get(ctx, "/")
	body := resp.Body()
	status := resp.Status()
	headers := resp.Headers()
	cookies := resp.Cookies()
	ok := resp.Ok()
	res := resp.Get()

	// Assert
	s.NoError(err)
	s.NotNil(resp)
	s.NotNil(body)
	s.Equal(http.StatusOK, status)
	s.Equal("application/json", headers.Get("Content-Type"))
	s.Equal(1, len(cookies))
	s.True(ok)
	s.NotNil(res)
}
