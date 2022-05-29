package controllermovie

import (
	"encoding/json"
	"fmt"
	"io"
	models "mrkresnofatih/golearning/gomuxapi/models"
	repositories "mrkresnofatih/golearning/gomuxapi/repositories"
	types "mrkresnofatih/golearning/gomuxapi/types"
	"net/http"

	mux "github.com/gorilla/mux"
)

func RegisterEndpointSaveMovie(r *mux.Router) {
	sr := r.PathPrefix("/saveMovie").Subrouter()
	sr.HandleFunc("/", SaveMovie).Methods(http.MethodPost)
}

var SaveMovie = types.BaseEndpoint(func(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

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

	movie := repositories.SaveRedisMovieById(movieToBeSaved.MovieId, *movieToBeSaved)

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
