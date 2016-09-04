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
			StartNodes:   []string{"52.79.49.39:7000", "52.79.49.39:7001", "52.79.49.39:7002", "52.79.49.39:7003", "52.79.49.39:7004", "52.79.49.39:7005"},
			ConnTimeout:  500 * time.Millisecond,
			ReadTimeout:  500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    600 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s \n", err.Error())
	}

	start := time.Now()

	insertNewCust(cluster)

	fmt.Printf("--Inserting %v user's info take %s...\n", maxUserCount, time.Since(start))
	fmt.Println("Press ANY key to continue...")
	fmt.Scanln()

	start = time.Now()
	getServicemgmtNo(cluster)
	fmt.Printf("**Retrieving %v user's info take %s...\n", maxUserCount, time.Since(start))

}

func insertNewCust(cluster *redis.Cluster) {
	baseMgmtNo := 1000000000
	baseExtrnid := 10000000

	for i := 0; i < maxUserCount; i++ {
		extrnid := "010" + strconv.Itoa(baseExtrnid)
		key := cePrefix + extrnid

		_, err := cluster.Do("HMSET", key, "serviceMgmtNo", baseMgmtNo, "extrnid", extrnid)

		if err != nil {
			log.Fatalf("cluster.Do error: %s\n", err.Error())
		}

		baseMgmtNo++
		baseExtrnid++
	}
}

func getServicemgmtNo(cluster *redis.Cluster) {
	baseMgmtNo := 1000000000
	baseExtrnid := 10000000

	for i := 0; i < maxUserCount; i++ {
		extrnid := "010" + strconv.Itoa(baseExtrnid)
		key := cePrefix + extrnid

		reply, err := redis.StringMap(cluster.Do("HGETALL", key))

		if err != nil {
			log.Fatalf("cluster.Do error: %s\n", err.Error())
		}

		fmt.Printf("custextrnid for %s is %s\n", key, reply)

		baseMgmtNo++
		baseExtrnid++
	}
}
