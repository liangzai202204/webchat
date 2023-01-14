package web

import "net/http"

var _ Server = &ServeHTTP{}

type ServeHTTP struct {
	routeRegister
	mdlS []Middleware
}

type ServeHTTPOption func(server *ServeHTTP)

func NewServeHTTP(opts ...ServeHTTPOption) *ServeHTTP {
	res := &ServeHTTP{
		routeRegister: newRouter(),
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
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
	n := s.findRoute(ctx.Req.URL.Path, ctx.Req.Method)
	if n == nil {
		ctx.RespW.WriteHeader(404)
		_, _ = ctx.RespW.Write([]byte("NOT FOUND"))
		return
	}
	// 寻找路由
	// 处理业务
	n.handle(ctx)
}

func (s *ServeHTTP) Post(path string, handleFunc HandleFunc) {
	s.addRoute(MethodPOST, path, handleFunc)
}
func (s *ServeHTTP) Get(path string, handleFunc HandleFunc) {
	s.addRoute(MethodGET, path, handleFunc)
}

//func ServerWithTemplateEngine(tplEngine TemplateEngine) HTTPServerOption {
//	return func(server *HTTPServer) {
//		server.tplEngine = tplEngine
//	}
//}

func ServerWithMiddleware(mdlS ...Middleware) ServeHTTPOption {
	return func(server *ServeHTTP) {
		server.mdlS = mdlS
	}
}
