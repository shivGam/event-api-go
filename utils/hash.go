package utils

import(
	"golang.org/x/crypto/bcrypt"
)

func HashPassWord(pass string) (string,error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(pass),14)
	return string(bytes),err
}

func ComparePasswords(hashedPass,pass string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPass),[]byte(pass))
	return err==nil
}