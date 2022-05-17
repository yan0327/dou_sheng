package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}
