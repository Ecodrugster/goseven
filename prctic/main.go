package main

import (
	"fmt"
	"project/mathutils" 
	"project/validator" 
	"project/calculator"
	"project/auth" 
)

func main() {
	fmt.Println("----- Math Utilities -----")
	fmt.Println("Add:", mathutils.Add(2, 3))
	fmt.Println("Subtract:", mathutils.Subtract(2, 3))
	fmt.Println("Multiply:", mathutils.Multiply(2, 3))
	fmt.Println("----- Validator -----")
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	fmt.Println("Возраст:", age)

	if err := validator.CheckAge(age); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Возраст корректный")
	}

	fmt.Println("----- Calculator -----")
	fmt.Print("Введите два числа: ")
	var a, b float64
	fmt.Scan(&a, &b)
	if _, err := calculator.Divide(a, b); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", calculator.Divide(a, b))
	}
	
}
