package main

import "fmt"

// 结构体是值类型
type person struct {
	name, city string
	age        int
}

func newPerson(name string, city string, age int) *person {
	return &person{ //指针类型可以减少内存开销
		name: name,
		city: city,
		age:  age,
	} //直接返回实例化后的结果
}

func main() {
	var p1 person
	var p2 = new(person)
	p3 := person{ //键值对初始化
		name: "man",
		city: "beijing",
		age:  20,
	}
	p4 := &person{
		name: "man",
		city: "beijing",
		age:  10,
	}
	p5 := newPerson("mankaiwen", "guangxi", 21)

	fmt.Println(p3, p4, p5)
	p1.name = "kevin"
	p1.city = "beijing"
	p1.age = 18
	p2.name = "kkvvm"
	(*p2).city = "jiangsu"
	(*p2).age = 18
	fmt.Printf("%#v", p2)
	fmt.Printf("\n")
	fmt.Printf("%#v", p1)
}
