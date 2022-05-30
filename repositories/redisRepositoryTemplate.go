package repositories

import (
	"context"
	"encoding/json"
	"time"

	utils "mrkresnofatih/golearning/gomuxapi/utils"

	redis "github.com/go-redis/redis/v8"
)

func getRedisClient() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rds
}

func GetUnmarshallableValueByKey(key string) ([]byte, error) {
	rds := getRedisClient()
	ctx := context.Background()
	val, err := rds.Get(ctx, key).Result()
	if err != nil {
		err = utils.WrapError("RedisClientErrorGet", err)
		return nil, err
	}
	return []byte(val), nil
}

func SaveMarshallableValueByKey(key string, value interface{}, ttl time.Duration) (*string, error) {
	rds := getRedisClient()
	ctx := context.Background()
	o, e := json.Marshal(value)
	if e != nil {
		e = utils.WrapError("JsonMarshalError", e)
		return nil, e
	}

	val, e := rds.Set(ctx, key, o, ttl).Result()
	if e != nil {
		e = utils.WrapError("RedisClientErrorGet", e)
		return nil, e
	}
	return &val, nil
}
