package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ErMamone/tWittor/middlew"
	"github.com/ErMamone/tWittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers port setter and listen*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/SingUp", middlew.CheckingDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckingDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
