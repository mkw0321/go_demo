package main

import "fmt"

func main() {
	mb := make(map[string]int)
	mb["23"] = 230231
	fmt.Println(mb)
	var ma map[string]int = make(map[string]int, 8) //map初始化
	//map中添加键值对
	ma["a"] = 100
	ma["b"] = 200
	ma["c"] = 300
	b := map[int]bool{ //声明并且初始化
		1: true,
		2: false,
	}

	c := make([]int, 0, 5)
	keys := make([]string, 0, 5)
	for k := range ma {
		keys = append(keys, k)
	}
	for k := range b {
		c = append(c, k)
	}
	for _, v := range c {
		fmt.Println(v)
	}
	for _, v := range ma {
		fmt.Println(v)
	}
	fmt.Println(b)

}
