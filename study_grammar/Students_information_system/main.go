package main

import (
	"fmt"
)

//requests1:Add students
//requests2:Edit students information
//requests3:show all students information

func showMenu() {
	fmt.Println("Welcome")
	fmt.Println("1.Add students")
	fmt.Println("2.Edit students information")
	fmt.Println("3.List students information")
	fmt.Println("4.Exit system")
}

func main() {

	for {
		showMenu() //print menu
		var input int
		fmt.Scanf("%d", &input)
		fmt.Printf("%d\n", input)
		switch input {
		case 1:
		case 2:
		case 3:
		case 4:

		}
	}
}
