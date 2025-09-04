package main

import (
	"fmt"
	"math"
)

func main() {
	const power = 2

	var userHeight float64
	var userWeight float64

	fmt.Println("**Калькулятор индекса веса**")
	fmt.Print("Введите свой рост в см:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг:")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight/100, power)

	fmt.Printf("Ваш индекс: %.0f\n", IMT)

}
