package controllermovie

import (
	"encoding/json"
	"fmt"
	repositories "mrkresnofatih/golearning/gomuxapi/repositories"
	types "mrkresnofatih/golearning/gomuxapi/types"
	utils "mrkresnofatih/golearning/gomuxapi/utils"
	"net/http"

	mux "github.com/gorilla/mux"
)

func RegisterEndpointGetMovieById(r *mux.Router) {
	sr := r.PathPrefix("/getMovieById").Subrouter()
	sr.HandleFunc("/{id}", GetMovieById).Methods(http.MethodGet)
}

var GetMovieById = types.BaseEndpoint(func(w http.ResponseWriter, r *http.Request) {
	movieId := mux.Vars(r)["id"]
	movie, err := repositories.GetRedisMovieById(movieId)
	if err != nil {
		err = utils.WrapError("GetRedisMovieById Error", err)
		utils.HandleErrorReturns(err, w)
		return
	}

	f, e := json.Marshal(movie)
	if e != nil {
		e = utils.WrapError("JsonMarshal Error", e)
		utils.HandleErrorReturns(e, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(f))
})
