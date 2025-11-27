package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	if n > 0 {

		if n%2 == 0 {
			fmt.Println("positive even")
		} else {
			fmt.Println("positive odd")
		}

	} else if n < 0 {

		if n%2 == 0 {
			fmt.Println("negative even")
		} else {
			fmt.Println("negative odd")
		}

	} else {
		fmt.Println("zero")
	}
}
