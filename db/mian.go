package main

import (
	"log"
	"step/db/leveldb"
)

func main() {
	//leveldb
	client, err := leveldb.NewClient("./db/leveldb/data/")
	if err != nil {
		log.Fatalln(err)
	}
	client.Demo()
}
