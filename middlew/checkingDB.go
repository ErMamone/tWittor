package middlew

import (
	"github.com/ErMamone/tWittor/bd"
	"net/http"

)

/*CheckingDB is a middleware for checking DB status*/
func CheckingDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ConnectionCheck() == 0 {
			http.Error(w, "Connection Lost", 500)
			return
		}
		next.ServeHTTP(w, r)

	}
}
