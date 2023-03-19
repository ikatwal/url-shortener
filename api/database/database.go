package database

import (
	"os"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var Context = context.Background()

func NewClient(numInstance int) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       numInstance,
	})
	return db
}
