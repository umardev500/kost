package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hashed password from a given password string.
//
// Takes in a string parameter `password` and returns a string `pass` and an error `err`.
func HashPassword(password string) (pass string, err error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

// ComparePassword compares a hashed password with a plaintext password and returns whether they match.
//
// hashedPass: the hashed password to compare.
// pass: the plaintext password to compare.
// success: a boolean indicating whether the passwords match.
func ComparePassword(hashedPass, pass string) (success bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
