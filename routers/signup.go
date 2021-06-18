package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/models"
)

/*SignUp is for a new user register  */
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error with data recived "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Entity without Email ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password length less than 6 characters ", 400)
		return
	}

	_, encontrado, _ := bd.CheckingIfUserExist(t.Email)

	if encontrado {
		http.Error(w, "This email is alredy used ", 400)
		return
	}

	_, status, err := bd.InsertNewUser(t)
	if err != nil {
		http.Error(w, "Error persisting or email has already exist "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Status problem ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
