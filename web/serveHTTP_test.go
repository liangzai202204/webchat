package web

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServeHTTP_Get(t *testing.T) {
	testadd := []struct {
		path   string
		method string
	}{
		{
			path:   "/",
			method: MethodPOST,
		},
	}
	r := newRouter()
	var f HandleFunc = func(ctx *Context) {}
	for _, ts := range testadd {
		r.addRoute(ts.method, ts.path, f)
	}

	testFind := []struct {
		name   string
		path   string
		method string
		node   *node
	}{
		{
			name:   "根节点",
			path:   "/",
			method: MethodPOST,
			node: &node{
				path:   "/",
				handle: f,
			},
		},
	}
	for _, tf := range testFind {
		t.Run(tf.name, func(t *testing.T) {
			n := r.findRoute(tf.path, tf.method)
			assert.Equal(t, tf.path, n.path)
		})
	}
}
