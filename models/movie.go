package models

import (
	json "encoding/json"
	utils "mrkresnofatih/golearning/gomuxapi/utils"
)

type Movie struct {
	MovieId     string   `json:"movieId"`
	Title       string   `json:"title"`
	Genres      []string `json:"genres"`
	Description string   `json:"description"`
	Year        int      `json:"year"`
	Watched     bool     `json:"watched"`
}

type MovieBuilder struct {
	Movie *Movie
}

func NewMovieBuilder() *MovieBuilder {
	return &MovieBuilder{Movie: &Movie{}}
}

func (m *MovieBuilder) SetID(id string) *MovieBuilder {
	m.Movie.MovieId = id
	return m
}

func (m *MovieBuilder) SetAutoID() *MovieBuilder {
	m.Movie.MovieId = utils.GeneratePrefixedId("mVe")
	return m
}

func (m *MovieBuilder) SetTitle(title string) *MovieBuilder {
	m.Movie.Title = title
	return m
}

func (m *MovieBuilder) SetDescription(desc string) *MovieBuilder {
	m.Movie.Description = desc
	return m
}

func (m *MovieBuilder) SetYear(year int) *MovieBuilder {
	m.Movie.Year = year
	return m
}

func (m *MovieBuilder) SetWatched(watched bool) *MovieBuilder {
	m.Movie.Watched = watched
	return m
}

func (m *MovieBuilder) SetGenres(genres []string) *MovieBuilder {
	m.Movie.Genres = genres
	return m
}

func (m *MovieBuilder) AddGenre(genre string) *MovieBuilder {
	if m.Movie.Genres != nil {
		m.Movie.Genres = append(m.Movie.Genres, genre)
	} else {
		m.Movie.Genres = []string{genre}
	}
	return m
}

func (m *MovieBuilder) Build() *Movie {
	return m.Movie
}

func (m *Movie) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
