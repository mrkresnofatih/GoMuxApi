package controllers

import (
	movieController "mrkresnofatih/golearning/gomuxapi/controllers/controllerMovie"

	mux "github.com/gorilla/mux"
)

func RegisterControllers(r *mux.Router) {
	rootRouter := r.PathPrefix("/api/v1").Subrouter()

	// Register Controllers
	movieController.RegisterMovieController(rootRouter)
}
