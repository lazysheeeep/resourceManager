package middleware

import "github.com/golang-jwt/jwt/v4"

type Option struct {
	Key string
	Val any
}

func WithOption(key string, val any) Option {
	return Option{
		Key: key,
		Val: val,
	}
}

func NewJwtToken(secretKey string, iat, seconds int64, opt ...Option) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat

	for _, v := range opt {
		claims[v.Key] = v.Val
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
