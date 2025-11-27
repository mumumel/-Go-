package main

import (
	"errors"
	"fmt"
)

func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	res1, res2 := a/b, a%b
	return res1, res2, nil
}

func main() {

	var a, b int
	fmt.Scan(&a, &b)
	ans1, ans2, err := divmod(a, b)
	if err != nil {
		fmt.Print("Ошибка: ", err)
	} else {
		fmt.Print(ans1, ans2)
	}
}
