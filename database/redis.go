package database

import (
	"blog/config"
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
	"time"
)

var redisClent *redis.Client

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Host + ":" + strconv.Itoa(config.Conf.Redis.Port),
		Password: "",
		DB:       config.Conf.Redis.DB,
	})
	redisClent = client

}

func GetRedis() *redis.Client {
	return redisClent
}

func CloseRedisClient() {
	redisClent.Close()
}

func redisTestConnect() {
	var key = "aa"
	var val = "bb"
	err := redisClent.Set(key, val, 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// key 命中错误 err 会有值
	redisVal, err := redisClent.Get(key).Result()
	if err != nil {
		panic(err)
	}

	if val != redisVal {
		fmt.Println("redis 操作异常")
		os.Exit(1)
	}

	redisClent.Del(key)
}
