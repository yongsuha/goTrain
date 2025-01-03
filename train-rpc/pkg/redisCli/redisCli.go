package redisCli

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Config struct {
	Cache struct {
		Type                 string `toml:"type"`
		Host                 string `toml:"host"`
		Port                 int    `toml:"port"`
		Password             string `toml:"password"`
		MaxIdleConnections   int    `toml:"max_idle_connections"`
		MaxActiveConnections int    `toml:"max_active_connections"`
	} `toml:"cache"`
}

var (
	redisCli *redis.Client
)

func GetRedis(db int) *redis.Client {
	if redisCli != nil {
		return redisCli
	}
	var config Config
	_, err := toml.DecodeFile("/Users/tal/go/goTrain/config/config.toml", &config)
	if err != nil {
		log.Fatalf("Failed to decode config.toml: %v", err)
	}

	// 建立 Redis 连接
	redisCli = redis.NewClient(&redis.Options{
		Addr:         config.Cache.Host + ":" + string(config.Cache.Port),
		Password:     config.Cache.Password,
		DB:           db, // 使用默认的数据库
		PoolSize:     config.Cache.MaxActiveConnections,
		MinIdleConns: config.Cache.MaxIdleConnections,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  100 * time.Second,
	})
	// 测试连接是否正常
	_, err = redisCli.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis")
	return redisCli
}
