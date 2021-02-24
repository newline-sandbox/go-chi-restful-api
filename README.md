# Go and `chi` RESTful API

This project demonstrates a simple RESTful API built with Go and [`chi`](https://github.com/go-chi/chi). This API provides the following endpoints:

* `GET /` - Verify whether or not the service is up and running ("health check"). Returns the "Hello World!" message 
* `GET /posts` - Retrieve a list of posts.
* `POST /posts` - Creates a post.
* `GET /posts/{id}` - Retrieve a single post identified by its `id`. 
* `PUT /posts/{id}` - Update a single post identified by its `id`. 
* `DELETE /posts/{id}` - Delete a single post identified by its `id`.

## Get Started

Install the dependencies...

```shell
$ make install_deps
```

...then run the service:

```shell
$ make run_service
```

![Running the API](https://www.dl.dropboxusercontent.com/s/092gazd3ahypf6s/Screen%20Shot%202021-02-24%20at%2012.41.04%20AM.png)

Feel free to clone this project and add more endpoints!