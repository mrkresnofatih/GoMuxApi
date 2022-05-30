package controllermovie

import (
	"encoding/json"
	"fmt"
	"io"
	models "mrkresnofatih/golearning/gomuxapi/models"
	repositories "mrkresnofatih/golearning/gomuxapi/repositories"
	types "mrkresnofatih/golearning/gomuxapi/types"
	utils "mrkresnofatih/golearning/gomuxapi/utils"
	"net/http"

	mux "github.com/gorilla/mux"
)

func RegisterEndpointSaveMovie(r *mux.Router) {
	sr := r.PathPrefix("/saveMovie").Subrouter()
	sr.HandleFunc("/", SaveMovie).Methods(http.MethodPost)
}

var SaveMovie = types.BaseEndpoint(func(w http.ResponseWriter, r *http.Request) {
	b, e := io.ReadAll(r.Body)
	defer r.Body.Close()
	if e != nil {
		e = utils.WrapError("IOReadAllError", e)
		utils.HandleErrorReturns(e, w)
		return
	}

	movieFromBody := models.Movie{}
	json.Unmarshal(b, &movieFromBody)

	movieToBeSaved := models.
		NewMovieBuilder().
		SetTitle(movieFromBody.Title).
		SetDescription(movieFromBody.Description).
		SetWatched(movieFromBody.Watched).
		SetYear(movieFromBody.Year).
		SetGenres(movieFromBody.Genres).
		SetAutoID().
		Build()

	movie, err := repositories.SaveRedisMovieById(movieToBeSaved.MovieId, *movieToBeSaved)
	if err != nil {
		err := utils.WrapError("SaveRedisMovieByIdError", err)
		utils.HandleErrorReturns(err, w)
		return
	}

	f, er := json.Marshal(movie)
	if er != nil {
		er = utils.WrapError("JsonMarshal Error", er)
		utils.HandleErrorReturns(er, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(f))
})
