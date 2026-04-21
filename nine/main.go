package main
import (
	"fmt"
	"errors"
	"strings"
)

type ValidationError struct{
	Field string
	Message string
}
func(v ValidationError) Error() string{
	return fmt.Sprintf("validation failed: %s %s", v.Field, v.Message)
}

func register(username string, email string, password string) error{
	if len(username) < 4{
		return ValidationError{
			Field: "username",
			Message: "is too short",
		}
	}
	if !strings.Contains(email, "@"){
		return ValidationError{
			Field: "email",
			Message: "does not contain @",
		}
	}
	if len(password) < 8{
		return ValidationError{
			Field: "password",
			Message: "is too short",
		}
	}
	hasDigit := false
	for i := 0; i < len(password); i++{
		if strings.Contains("0123456789", string(password[i])){
			hasDigit = true
			break
		}
	}
	if !hasDigit{
		return ValidationError{
			Field: "password",
			Message: "does not contain a digit",
		}
	}
	return nil
}


func main(){
	err := register("ab", "test@example.com", "password123")
	if err != nil{
		var validationErr ValidationError

		if errors.As(err, &validationErr){
			fmt.Println("Field:", validationErr.Field)
			fmt.Println("Message:", validationErr.Message)
			return
		}
		fmt.Println("Other error", err)
	}
}