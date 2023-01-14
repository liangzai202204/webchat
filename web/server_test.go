package web

import (
	"fmt"
	"log"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewServeHTTP()
	s.Get("/", func(ctx *Context) {
		fmt.Print("这是post", "\n")
		_, _ = ctx.RespW.Write([]byte(`haha`))
	})
	err := s.Start(":8080")
	if err != nil {
		log.Panicf("web:服务启动失败%t", err)
	}
}
