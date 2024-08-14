package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数sayhi,如果有第二个参数那么一定是error
	//如果有第二个参数那么一定是error
	sayHi := func(name string) (string, error) {
		return "Hello " + name, nil
	}
	//定义模板
	//解析模板
	t := template.New("hello.tmpl") //lianshi
	t.Funcs(template.FuncMap{
		"hello": sayHi,
	})
	_, err := t.ParseFiles("./gin_demo/demo6/hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	name := "xiaowangzi"
	//渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
