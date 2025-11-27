package main

import "fmt"

func main() {
	PI := 3.14159
	var r float64
	fmt.Scan(&r)
	area := PI * r * r
	fmt.Printf("Area = %.3f", area)
}
