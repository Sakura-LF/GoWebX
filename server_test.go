package GoWebX

import (
	"testing"
)

func TestServer(t *testing.T) {
	server := &HttpServer{}

	server.Get("/hello", func(ctx Context) {

	})
	server.Start(":8080")
}
