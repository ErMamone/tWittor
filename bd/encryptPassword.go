package bd

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword just for encrypt pwd :)*/
func EncryptPassword(pass string) (string, error) {
	costo := 8 //2 elevado al costo

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
