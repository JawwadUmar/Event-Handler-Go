package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(pasword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)

	return string(bytes), err
}
