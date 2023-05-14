package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("GoMirco")

func CheckToken(token string) error {
	if token == "" {
		return errors.New("need login")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := ParseToken(token)
	if err != nil {
		return err
	}
	return nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
