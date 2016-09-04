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

	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"52.79.49.39:7000", "52.79.49.39:7001", "52.79.49.39:7002", "52.79.49.39:7003", "52.79.49.39:7004", "52.79.49.39:7005"},
	})

	pong, err := cluster.Ping().Result()
	if err != nil {
		log.Panicf("Fail to Connect Redis cluster: %v, [%v]", pong, err)
	}
	fmt.Printf("Connect Redis cluster: %v, [%v]\n", pong, err)

	start := time.Now()

	//insertNewCust(cluster)

	fmt.Printf("--Inserting %v user's info take %s...\n", maxUserCount, time.Since(start))
	fmt.Println("Press ANY key to continue...")
	fmt.Scanln()

	start = time.Now()
	getServicemgmtNo(cluster)
	fmt.Printf("**Retrieving %v user's info take %s...\n", maxUserCount, time.Since(start))

}

func insertNewCust(cluster *redis.ClusterClient) {
	baseUserId := 10000000
	baseEuid := 1000000

	for i := 0; i < baseEuid; i++ {
		euid := baseEuid + i
		userid := baseUserId + 3*i
		key := euPrefix + strconv.Itoa(euid)

		hmap := map[string]string{
			"euid":   strconv.Itoa(euid),
			"userid": strconv.Itoa(userid),
		}

		res := cluster.HMSet(key, hmap)
		if res.Err() != nil {
			log.Printf("cluster.HMSET error: %s\n", res.Err())
		}
	}
}

func getServicemgmtNo(cluster *redis.ClusterClient) {
	baseEuid := 1000000

	for i := 0; i < baseEuid; i++ {
		euid := baseEuid + i
		key := euPrefix + strconv.Itoa(euid)

		euidUserid, err := cluster.HGetAll(key).Result()

		if err != nil {
			log.Fatalf("cluster.HGetAll error: %s\n", err)
		}

		fmt.Printf("custextrnid for %s is %v\n", key, euidUserid)
	}
}
