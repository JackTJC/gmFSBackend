package cache

import (
	"fmt"

	"github.com/JackTJC/gmFS_backend/config"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitCache() {
	conf := config.GetInstance()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Reids.Host, conf.Reids.Port),
		DB:       0,
		Password: conf.Reids.Passwd,
	})
	res, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
