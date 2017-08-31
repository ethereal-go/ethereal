package utils

import (
	"golang.org/x/crypto/bcrypt"
)

/**
/ Hash password..
*/
func HashPassword(password []byte, cost int) (hash string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword), err

	//hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(hashedPassword))
	//
	//// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match
}

/**
/ Compare Hash password..
*/
func CompareHashPassword(hash, password []byte) bool {
	if bcrypt.CompareHashAndPassword(hash, password) != nil {
		return false
	}
	return true
}
