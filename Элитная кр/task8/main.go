package main

import "fmt"

func main() {

	var N uint
	fmt.Scan(&N)

	i0 := uint(0)
	fmt.Print(0, 1, " ")
	for i := uint(1); i+i0 < N; {
		fmt.Printf("%v ", i0+i)
		i, i0 = i+i0, i
	}
	fmt.Println("- числа фибоначчи")

	Nn := N
	for j := uint(1); j < N; j++ {
		Nn += j
	}
	fmt.Printf("%v - Сумма чисел от 1 до N", Nn)
}
