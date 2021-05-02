package gohttpclient

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TestSuite struct {
	suite.Suite
	client Client
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.client = New("https://jsonplaceholder.typicode.com/")
}

func (s *TestSuite) Test_GetRequest_ReturnsSuccess() {
	request, err := s.client.Get("posts/10")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 10
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_GetRequestWith_ReturnsSuccess() {
	request, err := s.client.GetWith("posts", Todo{Id: 11})
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got []Todo
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.NotEmpty(s.T(), got)
}

func (s *TestSuite) Test_GetRequestWith_WhenRequestIsInvalid_ReturnsError() {
	request, err := s.client.GetWith("posts", "Lorem ipsum dolor")
	require.Nil(s.T(), request)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_PostRequest_ReturnsSuccess() {
	request, err := s.client.Post("posts")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 101
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PostRequestWith_ReturnsSuccess() {
	request, err := s.client.PostWith("posts", Todo{Id: 21})
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 101
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PostRequestWith_WhenRequestIsInvalid_ReturnsError() {
	invalidBody := make(chan int)
	request, err := s.client.PostWith("posts", &invalidBody)
	require.Nil(s.T(), request)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_PutRequest_ReturnsSuccess() {
	request, err := s.client.Put("posts/30")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 30
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PutRequestWith_ReturnsSuccess() {
	request, err := s.client.PutWith("posts/31", Todo{Id: 31})
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 31
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PRequestWith_WhenRequestIsInvalid_ReturnsError() {
	invalidBody := make(chan int)
	request, err := s.client.PutWith("posts", &invalidBody)
	require.Nil(s.T(), request)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_PatchRequest_ReturnsSuccess() {
	request, err := s.client.Patch("posts/40")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 40
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PatchRequestWith_ReturnsSuccess() {
	request, err := s.client.PatchWith("posts/41", Todo{Id: 41})
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	want := 41
	err = json.Unmarshal(response.Get().Body, &got)
	require.NoError(s.T(), err)

	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_PatchRequestWith_WhenRequestIsInvalid_ReturnsError() {
	invalidBody := make(chan int)
	request, err := s.client.PatchWith("posts", &invalidBody)
	require.Nil(s.T(), request)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_DeleteRequest_ReturnsSuccess() {
	request, err := s.client.Delete("posts/50")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	require.Equal(s.T(), http.StatusOK, response.Get().StatusCode)
}

func (s *TestSuite) Test_DeleteRequestWith_ReturnsSuccess() {
	request, err := s.client.DeleteWith("posts/51", Todo{Id: 51})
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	require.Equal(s.T(), http.StatusOK, response.Get().StatusCode)
}

func (s *TestSuite) Test_DeleteRequestWith_WhenRequestIsInvalid_ReturnsError() {
	invalidBody := make(chan int)
	request, err := s.client.DeleteWith("posts", &invalidBody)
	require.Nil(s.T(), request)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_DoFunction_WhenMakeRequest_ReturnsError() {
	var invalidUrl string
	httpClient := New(invalidUrl)
	request, err := httpClient.Get("posts/10")
	require.NoError(s.T(), err)

	response, err := httpClient.Do(request)
	require.Nil(s.T(), response)
	require.Error(s.T(), err)
}

func (s *TestSuite) Test_ToFunction_ReturnsSuccess() {
	request, err := s.client.Get("posts/10")
	require.NoError(s.T(), err)

	response, err := s.client.Do(request)
	require.NoError(s.T(), err)

	var got Todo
	response.To(&got)

	want := 10
	require.Equal(s.T(), want, got.Id)
}

func (s *TestSuite) Test_ToFunction_WhenBodyIsInvalid_ReturnsError() {
	httpClient := New("http://google.com/")
	request, err := httpClient.Get("qweqwe")
	require.NoError(s.T(), err)

	response, err := httpClient.Do(request)
	require.NoError(s.T(), err)

	var invalidModel chan int
	response.To(&invalidModel)
	require.Nil(s.T(), invalidModel)
}
