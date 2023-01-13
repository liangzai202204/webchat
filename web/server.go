package web

import "net/http"

// Server 服务接口实现 封装原有的http.Handler，可直接使用
// Start方法作为服务启动的唯一
// AddRoute方法是添加路由的入口
type Server interface {
	http.Handler
	Start(addr string) error
	addRoute(method string, addr string, handleFunc HandleFunc)
}
