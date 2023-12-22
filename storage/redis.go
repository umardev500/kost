package storage

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	db *redis.Client
}

func NewRedisStorage(db *redis.Client) fiber.Storage {
	return &redisStorage{
		db: db,
	}
}

// Close implements fiber.Storage.
func (r *redisStorage) Close() error {
	return r.db.Close()
}

// Delete implements fiber.Storage.
func (r *redisStorage) Delete(key string) error {
	if len(key) <= 0 {
		return nil
	}
	return r.db.Del(context.Background(), key).Err()
}

// Get implements fiber.Storage.
func (r *redisStorage) Get(key string) ([]byte, error) {
	if len(key) <= 0 {
		return nil, nil
	}
	val, err := r.db.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return val, err
}

// Reset implements fiber.Storage.
func (r *redisStorage) Reset() error {
	return r.db.FlushDB(context.Background()).Err()
}

// Set implements fiber.Storage.
func (r *redisStorage) Set(key string, val []byte, exp time.Duration) error {
	if len(key) <= 0 || len(val) <= 0 {
		return nil
	}
	return r.db.Set(context.Background(), key, val, exp).Err()
}
