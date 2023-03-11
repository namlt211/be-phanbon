package crypt

import (
	"green/config"
	"green/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func RegisterAccessToken(u *models.User) (string, error){

	expTime := time.Now().Add(time.Minute * 120)
	claims := &config.JWTClaim{
		Id : u.Id,
		UserName : u.UserName,
		RegisteredClaims : jwt.RegisteredClaims{
			Issuer: "green",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	jwtSecret := os.Getenv("JWT_ACCESS_KEY")
	signingKey := []byte(jwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func CreateRefreshToken(u *models.User) (string, error){
	expTime := time.Now().Add(time.Minute * 60)
	claims := &config.JWTClaim{
		Id : u.Id,
		UserName : u.UserName,
		RegisteredClaims : jwt.RegisteredClaims{
			Issuer: "green",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	jwtSecret := os.Getenv("JWT_REFRESH_KEY")
	signingKey := []byte(jwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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