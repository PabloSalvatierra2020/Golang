package main

import (
	"log"

	"github.com/PabloSalvatierra2020/Golang/bd"
	"github.com/PabloSalvatierra2020/Golang/handlers"
)

func main() {

	if bd.CheckedConnection() == 0 {
		log.Fatal("sin conexion a la Base de datos")
		return
	}
	handlers.Handlers()

}
