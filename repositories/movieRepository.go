package repositories

import (
	"encoding/json"
	models "mrkresnofatih/golearning/gomuxapi/models"
	utils "mrkresnofatih/golearning/gomuxapi/utils"
	"time"
)

func GetRedisMovieById(id string) (*models.Movie, error) {
	key := "MVE#" + id
	res, err := GetUnmarshallableValueByKey(key)
	if err != nil {
		err = utils.WrapError("GetUnmarshallableValueByKeyError", err)
		return nil, err
	}

	movie := models.Movie{}
	err = json.Unmarshal(res, &movie)
	if err != nil {
		err = utils.WrapError("JsonUnmarshallMovieError", err)
		return nil, err
	}
	return &movie, nil
}

func SaveRedisMovieById(id string, movie models.Movie) (*models.Movie, error) {
	key := "MVE#" + id
	_, err := SaveMarshallableValueByKey(key, movie, 1*time.Hour)
	if err != nil {
		err = utils.WrapError("SaveMarshallableValueByKeyError", err)
		return nil, err
	}
	return &movie, nil
}
