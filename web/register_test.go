package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

// 路由注册测试用例
func TestAddRoute(t *testing.T) {
	testAddCCase := []struct {
		path   string
		method string
	}{
		// 全静态path
		{
			path:   "/",
			method: http.MethodPost,
		},
		{
			path:   "/login",
			method: http.MethodPost,
		},
		{
			path:   "/register",
			method: http.MethodPost,
		},
		{
			path:   "/user/msg",
			method: http.MethodPost,
		},
	}
	r := newRouter()
	var mockHandler HandleFunc = func(ctx *Context) {}
	for _, tA := range testAddCCase {
		r.addRoute(tA.method, tA.path, mockHandler)
	}
	wantRes := &routeRegister{
		trees: map[string]*node{
			http.MethodPost: &node{
				path:   "/",
				handle: mockHandler,
				children: map[string]*node{
					"login": &node{
						path:   "login",
						handle: mockHandler,
					},
					"register": &node{
						path:   "register",
						handle: mockHandler,
					},
					"user": &node{
						path: "user",
						children: map[string]*node{
							"msg": &node{
								path:   "msg",
								handle: mockHandler,
							},
						},
					},
				},
			},
		},
	}
	msg, ok := wantRes.equal(&r)
	// r.equal(wantRouter)
	assert.True(t, ok, msg)

}

// 寻找节点测试用例
func TestFindRoute(t *testing.T) {
	// 先添加路由
	testAddCCase := []struct {
		path   string
		method string
	}{
		// 全静态path
		{
			path:   "/",
			method: http.MethodPost,
		},
		{
			path:   "/login",
			method: http.MethodPost,
		},
		{
			path:   "/register",
			method: http.MethodPost,
		},
		{
			path:   "/user/msg",
			method: http.MethodPost,
		},
	}
	r := newRouter()
	var mockHandler HandleFunc = func(ctx *Context) {}
	for _, tA := range testAddCCase {
		r.addRoute(tA.method, tA.path, mockHandler)
	}

	testFind := []struct {
		name   string
		path   string
		method string
		node   *node
	}{
		{
			name:   "/login",
			path:   "/login",
			method: http.MethodPost,
			node: &node{
				path:   "login",
				handle: mockHandler,
			},
		},
		{
			name:   "两层静态 handle",
			path:   "/user/msg",
			method: http.MethodPost,
			node: &node{
				path:   "msg",
				handle: mockHandler,
			},
		},
		{
			name:   "一层静态 no handle",
			path:   "/user",
			method: MethodPOST,
			node: &node{
				path: "user",
				children: map[string]*node{
					"msg": &node{
						path:   "msg",
						handle: mockHandler,
					},
				},
			},
		},
	}
	for _, tf := range testFind {
		t.Run(tf.name, func(t *testing.T) {
			n := r.findRoute(tf.path, tf.method)
			msg, ok := n.equal(tf.node)
			assert.True(t, ok, msg)
		})
	}
}

func (r *routeRegister) equal(y *routeRegister) (string, bool) {
	for k, v := range r.trees {
		dst, ok := y.trees[k]
		if !ok {
			return fmt.Sprintf("找不到对应的 http method"), false
		}
		msg, equal := v.equal(dst)
		if !equal {
			return msg, false
		}
	}
	return "", true
}
func (n *node) equal(y *node) (string, bool) {
	if n.path != y.path {
		return fmt.Sprintf("节点路径不匹配"), false
	}
	if len(n.children) != len(y.children) {
		return fmt.Sprintf("子节点数量不相等"), false
	}

	if n.starChild != nil {
		msg, ok := n.starChild.equal(y.starChild)
		if !ok {
			return msg, ok
		}
	}

	if n.paramChild != nil {
		msg, ok := n.paramChild.equal(y.paramChild)
		if !ok {
			return msg, ok
		}
	}

	// 比较 handler
	nHandler := reflect.ValueOf(n.handle)
	yHandler := reflect.ValueOf(y.handle)
	if nHandler != yHandler {
		return fmt.Sprintf("handler 不相等"), false
	}

	for path, c := range n.children {
		dst, ok := y.children[path]
		if !ok {
			return fmt.Sprintf("子节点 %s 不存在", path), false
		}
		msg, ok := c.equal(dst)
		if !ok {
			return msg, false
		}
	}
	return "", true
}
