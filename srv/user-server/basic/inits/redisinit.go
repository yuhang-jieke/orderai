package inits

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	conf := GetRedisConfigFromNacosOrLocal()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf["Host"], conf["Port"]),
		Password: conf["Password"].(string), // no password set
		DB:       conf["Database"].(int),    // use default DB
	})
	defer Rdb.Close()

	err := Rdb.Ping(ctx).Err()
	if err != nil {
		panic("redis连接失败")
	}
	fmt.Println("redis连接成功")
}
