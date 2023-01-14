package web

import (
	"fmt"
	"log"
	"strings"
)

type node struct {
	path   string
	node   *node
	handle HandleFunc
	// 子节点的 path => node
	children map[string]*node
	// todo
	starChild  *node
	paramChild *node
}

func (n *node) addChild(path string) *node {
	if path == "" {
		panic(fmt.Sprintf("web:路由注册的节点不能为空%s", path))
	}
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[path]
	if !ok {
		res = &node{
			path: path,
		}
		n.children[path] = res
	}
	return res
}

func (n *node) findChildOf(path string) *node {
	if n.children == nil {
		log.Print("web:没有子节点")
		return nil
	}
	root, ok := n.children[path]
	if !ok {
		return nil
	}
	return root
}

type routeRegister struct {
	trees map[string]*node
}

func newRouter() routeRegister {
	return routeRegister{
		trees: map[string]*node{},
	}
}

func (r *routeRegister) addRoute(method string, path string, handleFunc HandleFunc) {
	if path == "" {
		panic(fmt.Sprintf("web:路由注册路径不能为空%s", path))
	}
	root, ok := r.trees[method]
	if !ok {
		// 创建root，否则下面进入 path == "/" 直接退出，因为root==nil
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	if path == "/" {
		// 根节点重复注册
		if root.handle != nil {
			panic("web: 路由冲突，重复注册[/]")
		}
		root.handle = handleFunc
		//root.path = "/"
		return
	}
	pathS := strings.Split(path, "/")[1:]
	for _, p := range pathS {
		child := root.addChild(p)
		root = child
	}
	root.handle = handleFunc
	return
}
func (r *routeRegister) findRoute(path string, method string) *node {
	root, ok := r.trees[method]
	if !ok {
		return nil
	}
	if path == "/" {
		return root
	}
	paths := strings.Split(path, "/")[1:]
	for _, p := range paths {
		child := root.findChildOf(p)
		if child == nil {
			return nil
		}
		root = child
	}
	return root
}
