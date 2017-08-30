package ethereal

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type EtherealClaims struct {
	jwt.StandardClaims
}

// get key jwt
func JWTKEY() []byte {
	return []byte("AllYourBase")
}

// handler check error
func handlerErrorToken(token *jwt.Token, err error) (result bool, message error) {
	if token.Valid {
		return true, errors.New("You look nice today") // TODO add locale
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("That's not even a token")
			fmt.Println("")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return false, errors.New("Timing is everything")
		} else {
			return false, errors.New("Couldn't handle this token: " + err.Error())
		}
	} else {
		return false, errors.New("Couldn't handle this token: " + err.Error())
	}
	return
}
