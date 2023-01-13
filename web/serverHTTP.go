package web

import "net/http"

var _ Server = &ServeHTTP{}

type ServeHTTP struct {
	RouteRegister
}

// ServeHTTP 这是http请求入口
func (s *ServeHTTP) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 封装Context
	ctx := &Context{
		Req:   request,
		RespW: writer,
	}
	// 处理业务逻辑
	s.serve(ctx)
}

func (s *ServeHTTP) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *ServeHTTP) serve(ctx *Context) {
	s.findRoute(ctx.Req.URL.Path, ctx.Req.Method)
	// 寻找路由
	// 处理业务
}

func (s *ServeHTTP) Post(path string, handleFunc HandleFunc) {
	s.addRoute(MethodPOST, path, handleFunc)
}
func (s *ServeHTTP) Get(path string, handleFunc HandleFunc) {
	s.addRoute(MethodGET, path, handleFunc)
}
