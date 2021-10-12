package bd

import (
	"github.com/ErMamone/tWittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLoggin and check that login with DB*/
func TryLoggin(email string, password string) (models.User, bool) {
	usu, found, _ := CheckingIfUserExist(email)

	if found == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
