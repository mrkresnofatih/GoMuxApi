package controllermovie

import (
	"encoding/json"
	"fmt"
	repositories "mrkresnofatih/golearning/gomuxapi/repositories"
	types "mrkresnofatih/golearning/gomuxapi/types"
	"net/http"

	mux "github.com/gorilla/mux"
)

func RegisterEndpointGetMovieById(r *mux.Router) {
	sr := r.PathPrefix("/getMovieById").Subrouter()
	sr.HandleFunc("/{id}", GetMovieById).Methods(http.MethodGet)
}

var GetMovieById = types.BaseEndpoint(func(w http.ResponseWriter, r *http.Request) {
	movieId := mux.Vars(r)["id"]
	movie := repositories.GetRedisMovieById(movieId)

	f, e := json.Marshal(movie)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		f, er := json.Marshal(map[string]string{
			"Message": "Cannot Parse To Type",
			"Code":    "4001",
		})
		if er != nil {
		}
		fmt.Fprint(w, string(f))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(f))
})
