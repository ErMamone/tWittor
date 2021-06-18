package main

import (
	"log"

	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/handlers"
)

func main() {
	if bd.ConnectionCheck() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Handlers()
}
