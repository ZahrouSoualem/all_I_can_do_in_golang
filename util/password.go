package util

import bcrypt "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashpassword), nil
}

func CheckPassword(hashedpass string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
}
