package main

import (
	"flag"
	"log"

	"github.com/marcboeker/tentacle/server"
)

var (
	listenAddr, dbPath, testData *string
)

func init() {
	listenAddr = flag.String("listen", ":1337", "specify the interface and port to listen on")
	dbPath = flag.String("db", "peers.db", "path of the database file to store peers")
	testData = flag.String("test-data", "", "create testdata on startup")
	flag.Parse()
}

func main() {
	db, err := server.NewSQLiteDB(*dbPath)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(false)

	if testData != nil {
		insertTestData(db)
	}

	if err := server.StartAPIServer(*listenAddr, db); err != nil {
		log.Fatal(err)
	}
}

// Create 3 nodes and 2 subnets.
func insertTestData(db *server.DB) {
	node1 := server.Node{Name: "node1", Token: "token_1"}
	db.Create(&node1)
	node2 := server.Node{Name: "node2", Token: "token_2"}
	db.Create(&node2)
	node3 := server.Node{Name: "node3", Token: "token_3"}
	db.Create(&node3)

	subnet1 := server.Subnet{Name: "dev1", CIDR: "10.0.0.0/24"}
	db.Create(&subnet1)

	subnet2 := server.Subnet{Name: "dev2", CIDR: "10.0.1.0/24"}
	db.Create(&subnet2)

	db.Create(&server.SubnetNode{NodeID: node1.ID, SubnetID: subnet1.ID, IP: "10.0.0.1"})
	db.Create(&server.SubnetNode{NodeID: node1.ID, SubnetID: subnet2.ID, IP: "10.0.1.1"})

	db.Create(&server.SubnetNode{NodeID: node2.ID, SubnetID: subnet1.ID, IP: "10.0.0.2"})
	db.Create(&server.SubnetNode{NodeID: node2.ID, SubnetID: subnet2.ID, IP: "10.0.1.2"})

	db.Create(&server.SubnetNode{NodeID: node3.ID, SubnetID: subnet1.ID, IP: "10.0.0.3"})
	db.Create(&server.SubnetNode{NodeID: node3.ID, SubnetID: subnet2.ID, IP: "10.0.1.3"})
}
