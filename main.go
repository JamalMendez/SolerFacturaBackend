package main

import (
	"ggstudios/solerfacturabackend/db_connection"
	"ggstudios/solerfacturabackend/handlers"
	"ggstudios/solerfacturabackend/routes"
	"log"
	"net/http"
)

func main() {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	tipoProductoHandler := handlers.NewTipoProductoHandler()

	router := routes.InitRouter(tipoProductoHandler)

	port := ":8080"
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
