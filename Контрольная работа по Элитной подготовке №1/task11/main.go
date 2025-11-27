package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func reverse(arr []int) {

	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func main() {

	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	swap(&a, &b)
	fmt.Println(a, b)

	var N int

	fmt.Println("Введите размер массива:")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Вводите ваши числа:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	reverse(arr)
	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}
}
