package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	tmpl, err := template.ParseFiles("./gin_demo/demo4/hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	//渲染模板
	u := User{
		Name: "kkvvm",
		Age:  21,
	}
	tmpl.Execute(w, u)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}

}
