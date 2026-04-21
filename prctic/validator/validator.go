package validator
import "errors"

func CheckAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age < 18 {
		return errors.New("age must be at least 18")
	}
	return nil
}
