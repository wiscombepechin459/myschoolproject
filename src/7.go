package main

import "fmt"

func main() {
	// This is just an example function that prints out a random number between 1 and 100
	rand := generateRandomNumber(1, 100)
	fmt.Println("Generated number:", rand)
}

func generateRandomNumber(min int, max int) int {
	return min + (rand.Intn(max-min+1))
}
