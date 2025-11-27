package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func (c *Celsius) Add(d Celsius) {
	*c += d
}

func main() {

	var c float64
	fmt.Println("Введите число:")
	fmt.Scan(&c)

	cel := Celsius(c)
	fahr := cel.ToFahrenheit()

	fmt.Printf("Перевод %v в фаренгейты: %.2f\n", c, fahr)

	cel.Add(10)
	fmt.Printf("Перевод %v + 10 в фаренгейты: %.2f", c, cel.ToFahrenheit())
}
