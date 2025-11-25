package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)
	if score <= 100 && score >= 0 {
		switch score / 10 {
		case 10, 9:
			fmt.Println("A")
		case 8:
			fmt.Println("B")
		case 7:
			fmt.Println("C")
		case 6:
			fmt.Println("D")
		default:
			fmt.Println("F")
		}
	} else {
		fmt.Println("invalid")
	}
}
