package utils

import (
	"TreeHole/treehole_backend/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.JwtSecret)

type Claims struct {
	Username     string `json:"username"`
	EncrptedInfo string `json:"encryptedInfo"`
	jwt.StandardClaims
}

func GenerateToken(username, encryptedInfo string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)

	claims := Claims{
		username,
		encryptedInfo,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "treehole",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
