package main

import "fmt"

//方法定义实例

// define struct Person
type Person struct {
	name string
	age  int
}

// NewPerson is a constructor
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// change age
func (p *Person) SetAge(age int) {
	p.age = age
}

// define a method Dream   type is Person
func (p Person) Dream() {
	fmt.Printf("%s\n", p.name)
}

func main() {
	p1 := NewPerson("John", 20)
	p1.Dream()
	p1.SetAge(21)
	fmt.Println(p1.age)

}
