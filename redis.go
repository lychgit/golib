package golib

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"reflect"
	"time"
)

type RedisCacher interface {
	SetGlobalCache(key string, value interface{}, expiration time.Duration) error
	GetGlobalCache(key string) (string, error)
	Close() error
}

type RedisCache struct {
	rdb *redis.Client
}

var ctx = context.Background()

// 生成redis缓存操作接口
func NewRedisCacher (Options *redis.Options) (RedisCacher, error)  {
	rdb := redis.NewClient(Options)
	if rdb == nil {
		return nil, errors.New("redis服务器连接失败")
	}
	return &RedisCache{
		rdb: rdb,
	}, nil
}

func (rc *RedisCache) Close() error {
	return rc.rdb.Close()
}

func (rc *RedisCache) GetGlobalCache(key string) (string, error) {
	return rc.rdb.Get(ctx, key).Result()
}

func (rc *RedisCache) SetGlobalCache(key string, value interface{}, expiration time.Duration) error {
	var val string
	vType := reflect.TypeOf(value)
	if vType.Kind() == reflect.String {
		val = value.(string)
	} else {
		putByte, err := json.Marshal(value)
		if err != nil {
			return err
		}
		val = string(putByte)
	}
	// 放大时间倍数到秒级
	SetError := rc.rdb.Set(ctx, key, val, expiration * time.Second).Err()
	if SetError != nil {
		return SetError
	}
	return nil
}