package web

import (
	"fmt"
	"strings"
)

type node struct {
	path   string
	node   *node
	handle HandleFunc
}

type RouteRegister struct {
	trees map[string]*node
}

func newRouter() RouteRegister {
	return RouteRegister{
		trees: map[string]*node{},
	}
}

func (r *RouteRegister) addRoute(method string, path string, handleFunc HandleFunc) {
	root, ok := r.trees[method]
	if !ok {
		r.trees[method] = &node{
			path: "/",
		}
		return
	}
	if path == "/" {
		root.handle = handleFunc
	}
	root.node = &node{
		path:   path,
		handle: handleFunc,
	}
	return
}
func (r *RouteRegister) findRoute(path string, method string) *node {
	root, ok := r.trees[method]
	if !ok {
		return nil
	}
	if path == "/" {
		return root
	}
	paths := strings.Split(path, "/")[1:]
	fmt.Print(paths)
	return &node{}
}
