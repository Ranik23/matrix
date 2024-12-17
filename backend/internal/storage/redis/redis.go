package redis

import (
    "context"
    "errors"
    "github.com/go-redis/redis/v8"
)

type CacheStorage interface {
    Get(key string) (string, error)
    Set(key string, value string) error
    Delete(key string) error
}

type RedisStorage struct {
    client *redis.Client
}

func NewRedisStorage(addr string, password string, db int) *RedisStorage {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })
    return &RedisStorage{client: client}
}

func (r *RedisStorage) Get(key string) (string, error) {
    ctx := context.Background()
    val, err := r.client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", errors.New("key does not exist")
    } else if err != nil {
        return "", err
    }
    return val, nil
}

func (r *RedisStorage) Set(key string, value string) error {
    ctx := context.Background()
    return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisStorage) Delete(key string) error {
    ctx := context.Background()
    return r.client.Del(ctx, key).Err()
}