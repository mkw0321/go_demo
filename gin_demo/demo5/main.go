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
	t, err := template.ParseFiles("./gin_demo/demo5/hello.tmpl")
	if err != nil {
		fmt.Println("parse template fail")
	}
	u := User{
		Name: "jjj",
		Age:  22,
	}
	m1 := map[string]interface{}{
		"Name": "kkvvv",
		"Age":  18,
	}

	t.Execute(w, map[string]interface{}{
		"User": u,
		"m1":   m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
