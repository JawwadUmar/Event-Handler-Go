package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(pasword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)

	return string(bytes), err
}

func Validate(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
