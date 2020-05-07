package redis

import (
	"fmt"
	"time"

	"octopus/internal/cache"

	"github.com/go-redis/redis"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type redisCache struct {
	client            *redis.Client
	defaultExpiration time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewRedisCache returns a new RedisCache
func NewRedisCache(host, password string, defaultExpiration time.Duration) (cache.Cache, error) {
	// Create the client
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	// Check if the Redis responds
	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("Redis cache is unavailable")
	}

	return &redisCache{
		client:            client,
		defaultExpiration: defaultExpiration,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Set a value in Redis cache
func (c *redisCache) Set(key []byte, value interface{}, expires time.Duration) error {
	return c.client.Set(string(key), value, expires).Err()
}

// Get a value from Redis cache
func (c *redisCache) Get(key []byte) ([]byte, error) {
	return c.client.Get(string(key)).Bytes()
}
