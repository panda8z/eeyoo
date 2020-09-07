package utils

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

func ScryptPassword(password string) string {
	const keyLen = 10
	salt := []byte{53, 33, 26, 234, 123, 234, 42, 54}
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, keyLen)

	if err != nil {
		log.Fatal(err.Error())
		return password
	}
	ret := base64.StdEncoding.EncodeToString(dk)
	return ret
}
