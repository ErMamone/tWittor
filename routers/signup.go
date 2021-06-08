package routers

import (
	"encoding/json"
	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/models"
	"net/http"
)

/*SignUp is for a new user register  */
func SignUp(w http.ResponseWriter, r *http.Request){
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Entity without Email ", 400)
		return
	}

	if len(t.Password) <= 6 {
		http.Error(w, "Password length less than 6 characters ", 400)
		return
	}

	/*_, encontrado, _ := bd.CheckingIfUserExists(t.Email)

	if encontrado {
		http.Error(w, "This email is alredy used ", 400)
		return
	}
	*/
	_, status, err := bd.InsertNewUser(t)
	if err != nil {
		http.Error(w, "Error persisting "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Status problem ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
