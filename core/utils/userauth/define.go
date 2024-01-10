package userauth

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Username string
	Role     []string
	jwt.RegisteredClaims
}
