package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"webchat/ctrl"
	"webchat/web"
)

func main() {
	s := web.NewServeHTTP()
	s.Get("/", ctrl.Login)
	//RegisterView()
	err := s.Start(":8080")
	if err != nil {
		log.Panicf("web:服务启动失败%t", err)
	}
}
func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("vue/*/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}
