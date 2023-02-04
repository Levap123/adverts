package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var sign = []byte(viper.GetString("jwt_sign"))

type tokenClaims struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}

func GenerateJwt(UserID int, days int, TokenType string) (string, error) {
	fmt.Println(UserID, TokenType)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * time.Duration(days)).Unix(),
		},
		UserID,
		TokenType,
	})

	tokenString, err := token.SignedString(sign)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}
		return sign, nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", ErrInvalidClaims
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return 0, "", ErrExpired
	}
	return claims.UserID, claims.TokenType, nil
}
