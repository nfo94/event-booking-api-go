package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const dummykey = "dummykey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(dummykey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("not authorized")
		}

		return []byte(dummykey), nil
	})
	if err != nil {
		return 0, errors.New("not authorized")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("not authorized")
	}

	// Check type
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("not authorized")
	}

	// The int64 is read into float64
	// https://github.com/dgrijalva/jwt-go/issues/287
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
