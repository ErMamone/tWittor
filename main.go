package main

import (
	"github.com/ErMamone/tWittor/bd"
	"github.com/ErMamone/tWittor/handlers"
	"log"
)

func main(){
	if bd.ConnectionCheck()==0{
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Handlers()
}
