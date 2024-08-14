package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func add2(a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func main() {
	fmt.Println(add(20, 30))
}
