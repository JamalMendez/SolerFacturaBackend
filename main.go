package main

import (
	"ggstudios/solerfacturabackend/db_connection"
	"ggstudios/solerfacturabackend/routes"
	"log"
	"net/http"
)

func main() {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	router := routes.InitRouter()

	port := "0.0.0.0:8080"
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
