package routers

import (
	"errors"
	"strings"

	"github.com/ErMamone/bd"
	"github.com/ErMamone/models"
	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email devuelto del modelo*/
var Email string

/*ID devuelto del modelo*/
var IDUsuario string

/* ProcessToken procesamiento de token para extraer los valores que contiene*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("miMamaMeMima")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formate de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.CheckingIfUserExist(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}
