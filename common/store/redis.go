package store

<<<<<<< HEAD
import (
	"sync"
	"log"
	"github.com/go-redis/redis"
	"github.com/open-fightcoder/oj-vjudger/common/g"
)

var RedisClient *redis.Client
var once sync.Once

func InitRedis() {
	once.Do(func() {
		cfg := g.Conf().Redis
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.Database,
			PoolSize: cfg.PoolSize,
		})
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalln("fail to connect redis", g.Conf().Redis.Address, err)
	}
}

func CloseRedis() {
	RedisClient.Close()
=======
func InitRedis() {

}

func CloseRedis() {

>>>>>>> bd277c7ed08a2ecae5ddf41d9ae870e984e38ef2
}
