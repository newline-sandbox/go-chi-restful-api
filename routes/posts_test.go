package routes

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "log"
  "io/ioutil"
  "bytes"
)

type PostWithoutId struct {
  UserId int
  Title string
  Body string
}

type Post struct {
  Id int
  UserId int
  Title string
  Body string
}

type JsonPlaceholderMock struct {}

func (*JsonPlaceholderMock) GetPosts() (*http.Response, error) {
  mockedPosts := []Post{{
    Id: 1,
    UserId: 2,
    Title: "Hello World",
    Body: "Foo Bar",
  }}

  respBody, err := json.Marshal(mockedPosts)

  if err != nil {
    log.Panicf("Error reading mocked response data: %v", err)
  }

  return &http.Response{
    StatusCode: http.StatusOK,
    Body: ioutil.NopCloser(bytes.NewBuffer(respBody)),
  }, nil
}

func TestGetPostsHandler(t *testing.T) {
  GetPosts = (&JsonPlaceholderMock{}).GetPosts

  req, err := http.NewRequest("GET", "/posts", nil)

  if err != nil {
    t.Errorf("Error creating a new request: %v", err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(PostsResource{}.List)
  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
  }

  var posts []Post

  if err := json.NewDecoder(rr.Body).Decode(&posts); err != nil {
    t.Errorf("Error decoding response body: %v", err)
  }

  resultTotal := len(posts)
  expectedTotal := 1

  if resultTotal != expectedTotal {
    t.Errorf("Expected: %d. Got: %d.", expectedTotal, resultTotal)
  }
}

func BenchmarkGetPostsHandler(b *testing.B) {
  b.Run("Endpoint: GET /posts", func(b *testing.B) {
    GetPosts = (&JsonPlaceholderMock{}).GetPosts
    r, _ := http.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()
    handler := http.HandlerFunc(PostsResource{}.List)

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
      handler.ServeHTTP(w, r)
    }
  })
}

func BenchmarkGetPostsHandlerParallel(b *testing.B) {
  b.RunParallel(func(pb *testing.PB) {
    GetPosts = (&JsonPlaceholderMock{}).GetPosts
    r, _ := http.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()
    handler := http.HandlerFunc(PostsResource{}.List)

    b.ReportAllocs()
    b.ResetTimer()

    for pb.Next() {
      handler.ServeHTTP(w, r)
    }
  })
}
