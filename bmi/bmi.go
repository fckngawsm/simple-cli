package bmi

import (
	"fmt"
	"math"
)

func InitBMI() {
	var userHeight float64
	var userWeight float64

	fmt.Println("**Калькулятор индекса веса**")
	fmt.Print("Введите свой рост в см:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг:")
	fmt.Scan(&userWeight)

	IMT := calculateBMI(userWeight, userHeight)

	printBMI(IMT)

	minWeight, maxWeight := recommendWeight(userHeight)
	fmt.Printf("Подходящий вес для вас: %.1f - %.1f кг\n", minWeight, maxWeight)
}

func printBMI(bmi float64) {
	fmt.Printf("Ваш индекс массы тела: %.1f\n", bmi)
	if bmi < 16 {
		fmt.Println("Выраженный дефицит массы")
	} else if bmi < 16.9 {
		fmt.Println("Недостаток массы (умеренный)")
	} else if bmi < 18.4 {
		fmt.Println("Недостаток массы (легкий)")
	} else if bmi < 25 {
		fmt.Println("Норма (здоровый вес)")
	} else if bmi < 30 {
		fmt.Println("Избыточный вес")
	} else if bmi < 35 {
		fmt.Println("Ожирение I степени")
	} else if bmi < 40 {
		fmt.Println("Ожирение II степени")
	} else {
		fmt.Println("Ожирение III степени")
	}
}

func calculateBMI(userWeight, userHeight float64) float64 {
	const power = 2
	return userWeight / math.Pow(userHeight/100, power)
}

func recommendWeight(height float64) (float64, float64) {
	heightInMeters := height / 100
	minWeight := 18.5 * math.Pow(heightInMeters, 2)
	maxWeight := 24.9 * math.Pow(heightInMeters, 2)
	return minWeight, maxWeight
}
