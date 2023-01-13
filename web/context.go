package web

import "net/http"

type HandleFunc func(ctx *Context)

type Context struct {
	Req   *http.Request
	RespW http.ResponseWriter
}
