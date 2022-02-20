package main

import (
	"log"
	"step/db/myleveldb"
)

func main() {
	//leveldb
	client, err := myleveldb.NewClient("/Users/pzxy/WorkSpace/golang/tedge/core/edgex/cmd/tedge-resource/dbdata/992286/")
	if err != nil {
		log.Fatalln(err)
	}
	client.Demo()
}
