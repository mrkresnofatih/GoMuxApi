package repositories

import (
	"encoding/json"
	"log"
	models "mrkresnofatih/golearning/gomuxapi/models"
	"time"
)

func GetRedisMovieById(id string) *models.Movie {
	key := "MVE#" + id
	res := GetUnmarshallableValueByKey(key)
	if res == nil {
		return nil
	}

	movie := models.Movie{}
	err := json.Unmarshal(res, &movie)
	if err != nil {
		log.Panic(err)
	}

	return &movie
}

func SaveRedisMovieById(id string, movie models.Movie) *models.Movie {
	key := "MVE#" + id
	_, err := SaveMarshallableValueByKey(key, movie, 1*time.Hour)
	if err != nil {
		log.Panic(err)
	}

	return &movie
}
