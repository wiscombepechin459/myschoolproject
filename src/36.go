package main

import (
    "fmt"
    "math"
)

func main() {
    var height float64 = 10.5
    var radius float64 = math.Pi * (height / 2)
    fmt.Printf("The area of the circle is: %.3f\n", radius)
}
