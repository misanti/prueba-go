package main

import (
	"crudapirest/commons"
	"crudapirest/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	commons.Migrate()

	router:= mux.NewRouter()
	routes.SetRoutes(router)

	server := http.Server{
		Addr: ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}