package calculator

import "errors"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль не допустимо.")
	}
	return a / b, nil
}
