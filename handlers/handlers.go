package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/ErMamone/tWittor/routers"
	"github.com/ErMamone/tWittor/middlew"

	)

/*Handlers seteo de puerto y escucha del servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckingDB(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

