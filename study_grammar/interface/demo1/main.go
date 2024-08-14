package main

import "fmt"

type Bird struct{}

func (b Bird) Say() {
	fmt.Println("wangwangwang")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("wangwang")
}
func main() {
	b := Bird{}
	b.Say()
	d := Dog{}
	d.Say()

}
