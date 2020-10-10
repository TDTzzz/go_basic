package gee

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
}
