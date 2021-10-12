package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/jwt"
	"github.com/ErMamone/tWittor/models"
)

/*Login its just for login.... */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User and/or Password error "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "email required "+err.Error(), 400)
		return
	}

	document, exist := bd.TryLoggin(t.Email, t.Password)

	if exist == false {
		http.Error(w, "User and/or Password error ", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Token error "+err.Error(), 400)
		return
	}

	answ := models.AnswerLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answ)

	//Cookie save
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
