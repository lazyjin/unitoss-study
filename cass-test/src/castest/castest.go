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

var uimap map[string]interface{}

func main() {
	fmt.Println("Cassandra test with golang~!")

	cluster := gocql.NewCluster("52.79.148.85")
	cluster.Keyspace = "unitoss"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 3
	session, _ := cluster.CreateSession()
	defer session.Close()

	// var ui Userinfo
	uimap = make(map[string]interface{})

	// list all tweets
	iter := session.Query(`SELECT userid, userno, usedate, amount, charge, usetype FROM userinfo`).Iter()
	for iter.MapScan(uimap) {
		fmt.Println("USERINFO:", uimap["userid"], uimap["userno"], uimap["usedate"])
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
