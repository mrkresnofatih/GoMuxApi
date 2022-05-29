package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

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

func GetUnmarshallableValueByKey(key string) []byte {
	rds := getRedisClient()
	ctx := context.Background()
	val, err := rds.Get(ctx, key).Result()
	if err != nil {
		log.Panic(err)
		return nil
	}
	return []byte(val)
}

func SaveMarshallableValueByKey(key string, value interface{}, ttl time.Duration) (*string, error) {
	rds := getRedisClient()
	ctx := context.Background()
	o, e := json.Marshal(value)
	if e != nil {
		log.Panic(e)
		return nil, e
	}

	val, err := rds.Set(ctx, key, o, ttl).Result()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &val, nil
}
