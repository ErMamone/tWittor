package handlers

import (
	"github.com/ErMamone/tWittor/middlew"
	"github.com/ErMamone/tWittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
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

