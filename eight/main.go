package main 
import (
	"fmt"
	"errors"
)

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// 1
// func getDiscount(price float64) (float64, error) {
// 	if price < 0 {
// 		return 0, errors.New("price cannot be negative")
// 	}
// 	if price == 0 {
// 		return 0, errors.New("price cannot be zero")
// 	} 
// 	return price / 10.0, nil
// }
// 2

// func findStudentByName(name string) (string, error) {
// 	if name == "" {
// 		return "", errors.New("Имя не найдено")
// 	}
// 	return "Студент найден", nil
// }

// 3
func validateLogin(login string) (bool, error) {
	if login == "" {
		return false, errors.New("login cannot be empty")
	}
	if len(login) < 4 {
		return false, errors.New("login must be at least 5 characters long")
	}
	return true, nil
}

func main() {
	// 1
// 	result, err := getDiscount(100.0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }

// 2

// 	students := []string{"Alice", "Bob", "Charlie"}

// 	for _, student := range students {
// 		fmt.Println(student)
// 	}
// 	fmt.Println("End of program")

// 	result, err := findStudentByName("Alice")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// 3

	login := "alice"

	result, err := validateLogin(login)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}