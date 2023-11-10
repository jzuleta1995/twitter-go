package jwt

import (
	"errors"
	"strings"

	"twitter_go/models"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcessToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	myPass := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Token format is invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})

	if err == nil {

	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Invalid token")
	}

	return &claims, false, string(""), err
}
