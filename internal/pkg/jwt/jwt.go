package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"simple-demo/internal/pkg/global"
	"time"
)

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.JWTSetting.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	claims["issuer"] = global.JWTSetting.Issuer
	claims["expire"] = fmt.Sprintf("%d", time.Now().Unix()+int64(global.JWTSetting.Expire.Seconds()))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.JWTSetting.Secret))
}
