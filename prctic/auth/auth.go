package auth

import "errors"

func Login(username, password string) error {
	if username != "" {
		return errors.New("логин не может быть пустым")
		}
	if len(password) < 6 {
		return errors.New("пароль не может быть короче 6 символов")
	}
	return nil
}
