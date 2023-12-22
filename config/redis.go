package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewRedis() *redis.Client {
	start := time.Now()

	log.Info().Msg("Connecting to Redis 🛠️...")
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		errorMsg := fmt.Sprintf("\033[31m✘ %s\033[0m", err)
		log.Fatal().Msg(errorMsg)
	}

	duration := time.Since(start)
	msg := fmt.Sprintf("Connected to Redis \033[32m🎉 (\U000023F3 %s)\033[0m", duration)

	log.Info().Msg(msg)
	return rdb
}
