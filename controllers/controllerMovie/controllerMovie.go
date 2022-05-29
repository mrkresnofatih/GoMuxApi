package controllermovie

import (
	mux "github.com/gorilla/mux"
)

func RegisterMovieController(r *mux.Router) {
	s := r.PathPrefix("/movie").Subrouter()

	// Register Endpoints
	RegisterEndpointSaveMovie(s)
	RegisterEndpointGetMovieById(s)
}
