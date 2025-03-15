package main

import "fmt"

func main() {
	name := "Alice"
	age := 30

	if age > 26 {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println("Hi,", name)
	}
}
