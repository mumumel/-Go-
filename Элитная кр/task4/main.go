package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	if b != 0 {
		fmt.Printf("%.4f\n", a/b)
		fmt.Printf("%.4f", math.Mod(a, b))
	} else {
		fmt.Println("На ноль нельзя делить\n")
		fmt.Println("На ноль нельзя делить\n")
	}
}
