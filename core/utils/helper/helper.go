package helper

import "golang.org/x/crypto/bcrypt"

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
