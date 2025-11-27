package main

import (
	"fmt"
)

// Базовая пузырьковая сортировка с подсчетом перестановок
func bubbleSort(a *[5]int) int {
	n := 5
	swapCount := 0

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
				swapCount++
			}
		}
		if !swapped {
			break
		}
	}
	return swapCount
}

// Пузырьковая сортировка по убыванию
func bubbleSortDesc(a *[5]int) int {
	n := 5
	swapCount := 0

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
				swapCount++
			}
		}
		if !swapped {
			break
		}
	}
	return swapCount
}

// Шейкерная сортировка (сортировка перемешиванием)
func shakerSort(a *[5]int) int {
	n := 5
	swapCount := 0
	left := 0
	right := n - 1

	for left <= right {
		// Проход слева направо
		for i := left; i < right; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapCount++
			}
		}
		right--

		// Проход справа налево
		for i := right; i > left; i-- {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				swapCount++
			}
		}
		left++
	}
	return swapCount
}

func quickSort(a *[5]int) {
	sort(a, 0, len(a)-1)
}

func sort(a *[5]int, low, high int) {
	if low < high {
		pi := partition(a, low, high)
		sort(a, low, pi-1)
		sort(a, pi+1, high)
	}
}

func partition(a *[5]int, low, high int) int {
	pivot := a[high]
	i := low - 1

	for j := low; j < high; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func main() {
	var arr [5]int
	fmt.Println("Введите 5 чисел:")
	for i := 0; i < 5; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return
		}
	}

	// Копируем массив для демонстрации разных сортировок
	arr1 := arr
	arr2 := arr
	arr3 := arr
	arr4 := arr

	// Базовая пузырьковая сортировка (по возрастанию)
	swaps1 := bubbleSort(&arr1)
	fmt.Print("Сортировка по возрастанию: ")
	printArray(arr1)
	fmt.Printf("Перестановок: %d\n", swaps1)

	// Пузырьковая сортировка по убыванию
	swaps2 := bubbleSortDesc(&arr2)
	fmt.Print("Сортировка по убыванию: ")
	printArray(arr2)
	fmt.Printf("Перестановок: %d\n", swaps2)

	// Шейкерная сортировка
	swaps3 := shakerSort(&arr3)
	fmt.Print("Шейкерная сортировка: ")
	printArray(arr3)
	fmt.Printf("Перестановок: %d\n", swaps3)

	//Быстрая сортировка
	quickSort(&arr4)
	fmt.Print("Быстрая сортировка: ")
	printArray(arr4)

}

// Вспомогательная функция для вывода массива
func printArray(a [5]int) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
