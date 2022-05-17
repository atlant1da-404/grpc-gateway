package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var (
	rdb      *redis.Client
	ctx      = context.Background()
	noteCode = "note"
)

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

// this is implementation to get the last note id
func countToInt() int {

	count, err := rdb.Get(ctx, "count").Result()
	if err != nil {
		log.Fatalln(err.Error())
		return 0
	}

	currentId, _ := strconv.Atoi(count)
	return currentId
}

func incrementCounter() int {

	count := countToInt()

	_, err := rdb.Set(ctx, "count", count+1, 0).Result()
	if err != nil {
		log.Fatalln(err.Error())
		return 0
	}

	return count + 1
}
