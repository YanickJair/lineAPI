package utils

import "golang.org/x/crypto/bcrypt"

//HashPassword - hash plain text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err == nil {
		return string(bytes), nil
	}
	return "", err
}

//CheckPasswordHash - validate password
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}
