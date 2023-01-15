package ctrl

import (
	"text/template"
	"webchat/web"
)

func Login(ctx *web.Context) {
	//tpl := template.New("hello-world")
	tpl := template.New("hello-world")
	tp, _ := tpl.ParseFiles("vue/vue_test/login.html")
	err := tp.ExecuteTemplate(ctx.RespW, "/user/login.shtml", nil)
	if err != nil {
		return
	}
}
