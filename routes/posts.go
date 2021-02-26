package routes

import (
  "bytes"
  "context"
  "net/http"
  "io"
  "io/ioutil"

  "github.com/go-chi/chi"
)

var (
  GetPosts = (&JsonPlaceholder{}).GetPosts
  CreatePost = (&JsonPlaceholder{}).CreatePost
  GetPost = (&JsonPlaceholder{}).GetPost
  UpdatePost = (&JsonPlaceholder{}).UpdatePost
  DeletePost = (&JsonPlaceholder{}).DeletePost
)

type JsonPlaceholder struct {}

func (*JsonPlaceholder) GetPosts() (*http.Response, error) {
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  return resp, nil
}

func (*JsonPlaceholder) CreatePost(body io.ReadCloser) (*http.Response, error) {
  resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", body)

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  return resp, err
}

func (*JsonPlaceholder) GetPost(id string) (*http.Response, error) {
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  return resp, err
}

func (*JsonPlaceholder) UpdatePost(id string, body io.ReadCloser) (*http.Response, error) {
  client := &http.Client{}

  req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/" + id, body)
  req.Header.Add("Content-Type", "application/json")

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  resp, err := client.Do(req)

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  return resp, err
}

func (*JsonPlaceholder) DeletePost(id string) (*http.Response, error) {
  client := &http.Client{}

  req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/" + id, nil)

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  resp, err := client.Do(req)

  if err != nil {
    return &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer(nil)),
    }, err
  }

  return resp, err
}

type PostsResource struct {}

func (rs PostsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /posts - Read a list of posts.
	r.Post("/", rs.Create) // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
    r.Use(PostCtx)
		r.Get("/", rs.Get)       // GET /posts/{id} - Read a single post by :id.
		r.Put("/", rs.Update)    // PUT /posts/{id} - Update a single post by :id.
		r.Delete("/", rs.Delete) // DELETE /posts/{id} - Delete a single post by :id.
	})

	return r
}

// Request Handler - GET /posts - Read a list of posts.
func (rs PostsResource) List(w http.ResponseWriter, r *http.Request) {
  resp, err := GetPosts()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// Request Handler - POST /posts - Create a new post.
func (rs PostsResource) Create(w http.ResponseWriter, r *http.Request) {
  resp, err := CreatePost(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

// Request Handler - GET /posts/{id} - Read a single post by :id.
func (rs PostsResource) Get(w http.ResponseWriter, r *http.Request) {
  id := r.Context().Value("id").(string)
  resp, err := GetPost(id)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// Request Handler - PUT /posts/{id} - Update a single post by :id.
func (rs PostsResource) Update(w http.ResponseWriter, r *http.Request) {
  id := r.Context().Value("id").(string)
  resp, err := UpdatePost(id, r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// Request Handler - DELETE /posts/{id} - Delete a single post by :id.
func (rs PostsResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
  resp, err := DeletePost(id)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}