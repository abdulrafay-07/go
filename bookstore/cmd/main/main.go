package main

import (
	"log"
	"net/http"

	"github.com/abdulrafay-07/go/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	// register the routes
	routes.RegisterBookstoreRoutes(r)

	http.Handle("/", r)

	// create http server
	log.Fatal(http.ListenAndServe(":80", r))
}
