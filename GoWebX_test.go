package GoWebX

import "testing"

func TestGoWebX(t *testing.T) {
	server := &HttpServer{}

	server.Get("/hello", func(ctx Context) {

	})
	server.Start(":8080")
}
