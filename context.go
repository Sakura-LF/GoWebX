package GoWebX

import "net/http"

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter
}
