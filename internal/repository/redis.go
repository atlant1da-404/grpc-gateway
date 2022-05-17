package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdb *redis.Client

func NewRedis(socket string) {

	redisDB := redis.NewClient(&redis.Options{
		Addr:     socket,
		Password: "",
		DB:       0,
	})

	_, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	rdb = redisDB
}
