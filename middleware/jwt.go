package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	UserId int64
	jwt.StandardClaims
}

var signString = []byte("douyin")

func GenerateToken(userId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * 24 * time.Hour)
	issuer := "douyin_backend"
	claim := &UserClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(signString)

	if err != nil {
		return "", err
	}

	return token, nil

}

func ParseToken(tokenString string) (*UserClaims, bool) {
	tokenClaim, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signString, nil
	})
	if err != nil {
		fmt.Println("JWT: Token Parse Error: " + err.Error())
		return nil, false
	}

	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*UserClaims); ok && tokenClaim.Valid {
			return claims, true
		} else {
			return nil, false
		}
	}
	return nil, false
}
