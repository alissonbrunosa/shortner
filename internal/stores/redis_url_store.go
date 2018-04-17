package stores

import "github.com/go-redis/redis"

type RedisURLStore struct {
	Provider *redis.Client
}

func NewRedisURLStore() *RedisURLStore {
	return &RedisURLStore{
		Provider: redis.NewClient(&redis.Options{
			DB:       0,
			Addr:     "redis:6379",
			Password: "",
		}),
	}
}

func (this RedisURLStore) Create(key string, url string) error {
	err := this.Provider.Set(key, url, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (this RedisURLStore) Find(key string) (string, error) {
	return this.Provider.Get(key).Result()
}
