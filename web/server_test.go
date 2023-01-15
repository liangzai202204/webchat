package web

import (
	"log"
	"testing"
	"webchat/ctrl"
)

func TestServer(t *testing.T) {
	s := NewServeHTTP()
	s.Get("/", ctrl.Login)
	err := s.Start(":8080")
	if err != nil {
		log.Panicf("web:服务启动失败%t", err)
	}
}
