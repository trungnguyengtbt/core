package cryptor

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type Token struct {
	Username string
	jwt.StandardClaims
}

func GenerateTokenForUser(username string) string {
	tk := &Token{Username: username}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		panic("Error occured when generating token")
	}
	return tokenString
}
