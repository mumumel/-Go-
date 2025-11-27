package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	x = x + 3
	fmt.Printf("x = %v", x)

	x += 2
	fmt.Printf("x = %v", x)

	x *= 4
	fmt.Printf("x = %v", x)

	x -= 10
	fmt.Printf("x = %v", x)
}
