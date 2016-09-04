package common

import (
	"gopkg.in/redis.v4"
)

var RedisClust *redis.ClusterClient

func GetRedisCluster() *redis.ClusterClient {
	return RedisClust
}

func ConnectRedisCluster(addr []string) {
	RedisClust = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addr,
	})

	pong, err := RedisClust.Ping().Result()
	if err != nil {
		log.Panicf("Fail to Connect Redis cluster: %v, [%v]", pong, err)
	}

	log.Info("Successfully connect to Redis cluster...")
}
