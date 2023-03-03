package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateBearerToken(formula string, secret string, timeout int) (*string, *error) {

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["client"] = formula
	sign.Claims = claims

	token, err := sign.SignedString([]byte(secret))
	if err != nil {
		return nil, &err
	}

	bearerToken := "Bearer " + token

	return &bearerToken, nil
}

func ClaimJwt(bearerToken string, secret string) *error {

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return &err
	}

	if token == nil {
		newError := errors.New("token is nil")
		return &newError
	}

	return nil
}
