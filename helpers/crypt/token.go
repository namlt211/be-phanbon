package crypt

import (
	"github.com/golang-jwt/jwt/v4"
)

func RegisterAccessToken(secret string, c jwt.Claims) (string, error){
	signingKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParserToken(tokenString string, key string) (jwt.MapClaims, error){
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error){
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}