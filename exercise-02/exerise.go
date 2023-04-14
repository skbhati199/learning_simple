package main

import "fmt"

func main() {
	pi := 3.14

	radius := 5

	circumference := 2 * pi * float64(radius)
	area := pi * float64(radius) * float64(radius)
	// fmt.Println(circumference)
	fmt.Printf("%.2f\n", circumference)
	fmt.Printf("%.2f\n", area)
}