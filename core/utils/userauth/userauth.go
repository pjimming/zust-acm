package userauth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pjimming/zustacm/core/constant"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	return token.SigningString()
}
