package config

import "github.com/golang-jwt/jwt/v4"

var JWT_ACCESS_KEY = []byte("ashdjqY#9283409bsdklkg8hda01")
var JWT_REFRESH_KEY = []byte("ashdjqY#9283409bsdklkdsaasdw")


type JWTClaim struct {
	Id int64
	UserName string
	jwt.RegisteredClaims
}
