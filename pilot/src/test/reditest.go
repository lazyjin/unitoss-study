package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"log"
	"strconv"
	"time"
)

const maxUserCount = 100000
const euPrefix = "euid:"

func main() {
	fmt.Println("Redis test with golang~!")

	// cluster, err := redis.NewCluster(
	// 	&redis.Options{
	// 		StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
	// 		ConnTimeout:  50 * time.Millisecond,
	// 		ReadTimeout:  50 * time.Millisecond,
	// 		WriteTimeout: 50 * time.Millisecond,
	// 		KeepAlive:    16,
	// 		AliveTime:    60 * time.Second,
	// 	})
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"52.79.49.39:7000", "52.79.49.39:7001", "52.79.49.39:7002", "52.79.49.39:7003", "52.79.49.39:7004", "52.79.49.39:7005"},
	})

	pong, err := cluster.Ping().Result()
	if err != nil {
		log.Panicf("Fail to Connect Redis cluster: %v, [%v]", pong, err)
	}
	fmt.Printf("Connect Redis cluster: %v, [%v]\n", pong, err)

	// log.Debug("Successfully connect to Redis cluster...")

	start := time.Now()

	insertNewCust(cluster)

	fmt.Printf("--Inserting %v user's info take %s...\n", maxUserCount, time.Since(start))
	fmt.Println("Press ANY key to continue...")
	fmt.Scanln()

	start = time.Now()
	// getServicemgmtNo(cluster)
	fmt.Printf("**Retrieving %v user's info take %s...\n", maxUserCount, time.Since(start))

}

func insertNewCust(cluster *redis.ClusterClient) {
	baseUserId := 10000000
	baseEuid := 1000000

	for i := 0; i < baseEuid*2; i++ {
		euid := baseEuid + i
		userid := baseUserId + 3*i
		key := euPrefix + strconv.Itoa(baseEuid)

		hmap := map[string]string{
			"euid":   strconv.Itoa(euid),
			"userid": strconv.Itoa(userid),
		}

		res := cluster.HMSet(key, hmap)
		if res.Err() != nil {
			log.Fatalf("cluster.HMSET error: %s\n", res.Err())
		}
	}
}

// func getServicemgmtNo(cluster *redis.ClusterClient) {
// 	baseMgmtNo := 1000000000
// 	baseExtrnid := 10000000

// 	for i := 0; i < maxUserCount; i++ {
// 		extrnid := "010" + strconv.Itoa(baseExtrnid)
// 		key := cePrefix + extrnid

// 		reply, err := redis.StringMap(cluster.Do("HGETALL", key))

// 		if err != nil {
// 			log.Fatalf("cluster.Do error: %s\n", err.Error())
// 		}

// 		fmt.Printf("custextrnid for %s is %s\n", key, reply)

// 		baseMgmtNo++
// 		baseExtrnid++
// 	}
// }
