package cache

import (
	"fmt"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/config"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init()  {
	redisConfig := config.GetAppConfig().Application.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Database,  // use default DB
		DialTimeout: time.Duration(redisConfig.Timeout),
		PoolSize: redisConfig.PoolConfig.MaxActive,
		MinIdleConns: redisConfig.PoolConfig.MinIdle,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetClient() *redis.Client {
	return client
}
