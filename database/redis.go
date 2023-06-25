package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ConnectRedisDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-10351.c114.us-east-1-4.ec2.cloud.redislabs.com:10351",
		Password: "QIqrMPlHJJ8UuhFjr936khTiJwPE3ChP",
		DB:       0,
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}

	fmt.Println("Connected to Redis! 🎉")

}
