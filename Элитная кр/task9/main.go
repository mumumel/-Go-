package main

import "fmt"

func sum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func main() {
	var N int
	fmt.Println("Сколько чисел вы хотите ввести?")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Вводите ваши числа:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	result := sum(arr)
	fmt.Printf("Сумма чисел: %v\n", result)
}
