package main

//简单监听一个网站

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./gin_demo/demo2/hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil) //监听函数
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}

}
