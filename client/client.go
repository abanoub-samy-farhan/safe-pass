package client

import (
	"github.com/redis/go-redis/v9"	
)

func InitiateClient(database int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       database,
	})
	
	return client
}