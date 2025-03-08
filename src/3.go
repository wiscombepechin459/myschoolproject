package main

import "fmt"

func main() {
	name := "John Doe"
	age := 30
	isEmployee := true

	fmt.Println("Hello, my name is", name, "and I am", age, "years old.")
	if isEmployee {
		fmt.Println("I am an employee.")
	} else {
		fmt.Println("I am not an employee.")
	}
}
