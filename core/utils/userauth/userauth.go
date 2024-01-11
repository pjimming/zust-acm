package userauth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pjimming/zustacm/core/constant"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenCtxKey = "token"
	userCtxKey  = "user"
)

func EncryptPwd(plainPwd string) (string, error) {
	cipherPwd, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(cipherPwd), nil
}

func CheckPwd(dbPwd, inputPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(inputPwd))
}

func GenToken(username string) (string, error) {
	uc := &UserClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.TokenExpire)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(constant.JwtKey))
}

func GenTokenKey(token string) string {
	return constant.RedisTokenPrefix + "$" + token
}

func ParseToken(token string) (uc *UserClaim, err error) {
	uc = &UserClaim{}
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JwtKey), nil
	})

	if err != nil {
		return nil, errorx.ErrorTokenInvalid()
	}

	if !claims.Valid {
		return nil, errorx.ErrorTokenInvalid()
	}
	return
}

func CtxWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtxKey, token)
}

func CtxWithUser(ctx context.Context, user *UserClaim) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}
