package main

import (
	"fmt"
	"github.com/chasex/redis-go-cluster"
	"log"
	"strconv"
	"time"
)

const maxUserCount = 100000
const cePrefix = "cust"

func main() {
	fmt.Println("Redis test with golang~!")

	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
	}

	start := time.Now()

	insertNewCust(cluster)

	fmt.Println("Inserting %v user's info take %v...", time.Since(start))

	getServicemgmtNo(cluster)

}

func insertNewCust(cluster *redis.Cluster) {
	baseMgmtNo := 1000000000
	baseExtrnid := 10000000

	for i := 0; i < maxUserCount; i++ {
		extrnid := "010" + strconv.Itoa(baseExtrnid)
		key := cePrefix + extrnid

		_, err := cluster.Do("HMSET", key, "serviceMgmtNo", baseMgmtNo, "extrnid", extrnid)

		if err != nil {
			log.Fatalf("cluster.Do error: %s", err.Error())
		}

		baseMgmtNo++
		baseExtrnid++
	}
}

func getServicemgmtNo(cluster *redis.Cluster) {
	key := cePrefix + "010" + strconv.Itoa(10000000)
	reply, err := redis.StringMap(cluster.Do("HGETALL", key))

	if err != nil {
		log.Fatalf("redis.StringMap error: %s", err.Error())
	}

	fmt.Println("custextrnid for %v is %v", key, reply)

	key = cePrefix + "010" + strconv.Itoa(10000001)
	reply, err = redis.StringMap(cluster.Do("HGETALL", key))

	if err != nil {
		log.Fatalf("redis.StringMap error: %s", err.Error())
	}

	fmt.Println("custextrnid for %v is %v", key, reply)

	key = cePrefix + "010" + strconv.Itoa(10010000)
	reply, err = redis.StringMap(cluster.Do("HGETALL", key))

	if err != nil {
		log.Fatalf("redis.StringMap error: %s", err.Error())
	}

	fmt.Println("custextrnid for %v is %v", key, reply)
}
