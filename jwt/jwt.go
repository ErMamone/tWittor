package jwt

import (
	"time"

	"github.com/ErMamone/tWittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerateJWT generate an ecrypt with JWT*/
func GenerateJWT(t models.User) (string, error) {

	myPass := []byte("miMamaMeMima")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellido":        t.Apellido,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicaion":        t.Ubicaion,
		"sitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPass)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
