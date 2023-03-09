package cache

import (
	"encoding/json"
	"gorillamux/pkg/common/models"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// Making a new redis instance
func NewRedisCache(host string, db int, exp time.Duration) ProdCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}

}

// Generating a redis connection
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(ctx context.Context, key string, value *models.Product) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, json, cache.expires*time.Second*15)
}

func (cache *redisCache) Get(ctx context.Context, key string) *models.Product {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	prod := models.Product{}
	err = json.Unmarshal([]byte(val), &prod)
	if err != nil {
		panic(err)
	}
	return &prod
}
