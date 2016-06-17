package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

// type Userinfo struct {
// 	userid  string
// 	userno  string
// 	usedate string
// 	amount  int
// 	charge  int
// 	usetype string
// }

const clusterIP = "127.0.0.1"

var uimap map[string]interface{}
var uimap2 map[string]interface{}
var uimap2arr []map[string]interface{}

func main() {
	fmt.Println("***Cassandra test with golang~!****")

	cluster := gocql.NewCluster(clusterIP)
	cluster.Keyspace = "unitoss"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 3 // should be set to 3 or 4
	session, _ := cluster.CreateSession()
	defer session.Close()

	// get normal table date
	fmt.Print("Get userinfo using select iterator!!!\n\n")

	uimap = make(map[string]interface{})

	iter := session.Query(`SELECT userid, userno, usedate, amount, charge, usetype FROM userinfo`).Iter()

	fmt.Println("----------------------------------------------------------")
	for iter.MapScan(uimap) {
		fmt.Printf("USERINFO => userid:[%s], userno[%s], usedate:[%s], amount:[%d], charge:[%d], usetype:[%s]\n",
			uimap["userid"], uimap["userno"], uimap["usedate"], uimap["amount"], uimap["charge"], uimap["usetype"])
	}
	fmt.Println("----------------------------------------------------------")
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// get table include collection data type
	fmt.Print("Get userinfo2 using select iterator!!!\n\n")

	uimap2 = make(map[string]interface{})
	uimap2arr = make([]map[string]interface{}, 0, 5)

	iter = session.Query(`SELECT userid, username, age, gender, address FROM userinfo2`).Iter()

	fmt.Println("----------------------------------------------------------")
	for iter.MapScan(uimap2) {
		uimap2arr = append(uimap2arr, uimap2)

		addrMap := uimap2["address"]
		fmt.Println("USERINFO:",
			uimap2["userid"], uimap2["username"], uimap2["age"], uimap2["gender"], uimap2["address"])

		addr, ok := addrMap.(map[string]string)
		fmt.Println("***address key, value is ", addr, ok)

		if ok {
			fmt.Printf("\nADDRESS: city:[%s], street:[%s], zip:[%s]\n",
				addr["city"], addr["street"], addr["zip"])
		}

		uimap2 = make(map[string]interface{})
	}
	fmt.Println("----------------------------------------------------------")

	fmt.Println("uimap2arr: ", uimap2arr)

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
