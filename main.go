package main

import (
	"log"
	"net/http"

	controllers "mrkresnofatih/golearning/gomuxapi/controllers"

	mux "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	controllers.RegisterControllers(r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
