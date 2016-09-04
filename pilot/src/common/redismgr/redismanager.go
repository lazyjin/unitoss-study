package redismgr

import (
	"common/clog"
	"gopkg.in/redis.v4"
)

var RedisClust *redis.ClusterClient

var log = clog.GetLogger()

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
