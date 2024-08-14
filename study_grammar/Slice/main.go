package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [...]int{3, 1, 2}
	sort.Ints(a[0:2])
	fmt.Println(a)
}
