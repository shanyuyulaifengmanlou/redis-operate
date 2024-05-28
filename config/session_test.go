package config

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"testing"
)

func TestSession(t *testing.T) {
	ctx := context.Background()

	// 创建一个 redis 客户端。
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", DefaultRedisEndpoint, DefaultRedisPort), // Redis 地址
		Password: "",                                                           // Redis 密码，如果没有设置则为空
		DB:       0,                                                            // 使用默认的数据库
	})

	// 设置一个 key-value 对
	err := rdb.Set(ctx, "key", "value2", 0).Err()
	if err != nil {
		panic(err)
	}
}
