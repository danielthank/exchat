package infra

import (
	"log"

	"github.com/danielthank/redisstore"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
)

type RedisHandler struct {
	Client *redis.Client
	Store  *redisstore.RedisStore
}

func NewRedisHandler() *RedisHandler {
	redisClient := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	redisStore, err := redisstore.NewRedisStore(redisClient)
	if err != nil {
		log.Fatal("failed to create redis store")
	}
	redisStore.KeyPrefix("session_")
	redisStore.Options(sessions.Options{
		MaxAge: 86400 * 7, // 7 days
	})

	return &RedisHandler{
		Client: redisClient,
		Store:  redisStore,
	}
}
