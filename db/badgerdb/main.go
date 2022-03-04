package main

import (
	"fmt"
	badger "github.com/dgraph-io/badger/v3"
	"log"
	"net/http"
	_ "net/http/pprof"
	"step/db/badgerdb/mybadger"
	"strconv"
)

var ss = "Google leveldb简介 - 简书https://www.jianshu.com › ...\n5、支持事务机制。 6、支持数据库快照功能。 7、支持前向和后向遍历数据。 8、默认线程安全（除了WriteBatch等 ...\n\nLevelDB 入门—— 全面了解LevelDB 的功能特性 - 掘金https://juejin.cn › 后端\n2019年1月3日 — 为此LevelDB 提供了批处理功能，批处理操作就好比事务，LevelDB 确保这一些列写操作的原子性执行，要么全部生效要么完全不生效。 class WriteBatch { void ...\n\nLevelDB源码解析5. 数据完整性 - 知乎专栏https://zhuanlan.zhihu.com › ...\n2018年3月17日 — WAL LOG文件，或者说binlog。是用来支持事务完整性的。 $ cat 000003.log JS )KeyNameExample ValueExample. 对于WAL ...\n\n第四十章LevelDB与BoltDB - 《Go语言四十二章经》"

func main() {
	http.ListenAndServe("0.0.0.0:6061", nil)

	client, err := mybadger.NewClient("./db/badgerdb/data/")
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 10000000; i++ {
		err := client.PutString("a"+strconv.Itoa(i), ss)
		if err != nil {
			log.Fatalln(err)
			return
		}
		//fmt.Println("插入", i)
	}
}

func demo() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("./db/badgerdb/data/"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			log.Fatal(err)
		}

		var valNot, valCopy []byte
		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		// DO NOT access val here. It is the most common cause of bugs.
		fmt.Printf("NEVER do this. %s\n", valNot)

		// You must copy it to use it outside item.Value(...).
		fmt.Printf("The answer is: %s\n", valCopy)

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The answer is: %s\n", valCopy)

		return nil
	})

	defer db.Close()
}
