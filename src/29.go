package main

import "fmt"

func main() {
    var score int
    fmt.Println("Enter your name:")
    fmt.Scan(&score)
    fmt.Printf("Welcome, %s! Your current score is: %d\n", score, score)
}
