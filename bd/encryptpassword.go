package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(psw string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), cost)

	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
