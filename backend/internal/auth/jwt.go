package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("curly-secret")

func GenerateToken(
	userID uint,
	username string,
) (string, error) {

	claims := jwt.MapClaims{

		"user_id": userID,

		"username": username,

		"exp": time.Now().
			Add(
				time.Hour * 24,
			).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(secret)

}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			// 检查签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("签名算法错误")

			}

			return secret, nil
		},
	)

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {

		return nil, fmt.Errorf("token无效")

	}

	return claims, nil
}
