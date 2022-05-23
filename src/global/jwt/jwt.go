package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	issuer := "frank"
	claims := Claims{
		Id:   10001,
		Name: "wangruiyu",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*Claims, bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, false, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, true, nil
		}
	}
	return nil, false, err
}
