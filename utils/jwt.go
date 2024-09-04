/*
 * @Author: lihuan
 * @Date: 2024-09-01 20:42:11
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-04 20:53:59
 * @Email: 17719495105@163.com
 */
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	ID       int64  `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int64, username string, secret []byte, exp int) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(exp) * time.Hour)

	claims := Claims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "smartgreeting-mini-go",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(secret)

	return token, err
}

func ParseToken(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
