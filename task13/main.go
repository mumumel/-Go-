package main

import "fmt"

func Stats(arr []int) (sum int, maxIdx int, maxVal int, positives []int, average float64) {
	if len(arr) == 0 {
		return 0, 0, 0, []int{}, 0
	}

	// Сумма и максимальный элемент
	maxVal = arr[0]
	maxIdx = 0
	for i, v := range arr {
		sum += v
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}

	// Положительные элементы
	for _, v := range arr {
		if v > 0 {
			positives = append(positives, v)
		}
	}

	// Среднее значение
	average = float64(sum) / float64(len(arr))

	return sum, maxIdx, maxVal, positives, average
}

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}

	// Использование функции Stats
	sum, maxIdx, maxVal, positives, avg := Stats(a[:])

	// Вывод результатов
	fmt.Println(sum)
	fmt.Printf("%d %d\n", maxIdx, maxVal)

	// Вывод положительных элементов через пробел
	if len(positives) > 0 {
		for i, v := range positives {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	} else {
		fmt.Println()
	}

	// Дополнительно: вывод среднего с 2 знаками
	fmt.Printf("%.2f\n", avg)
}
