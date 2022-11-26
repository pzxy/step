package main

import (
	"log"
	_ "net/http/pprof"
	"step/db/leveldb/myleveldb"
)

var ss = "Google leveldb简介 - 简书https://www.jianshu.com › ...\n5、支持事务机制。 6、支持数据库快照功能。 7、支持前向和后向遍历数据。 8、默认线程安全（除了WriteBatch等 ...\n\nLevelDB 入门—— 全面了解LevelDB 的功能特性 - 掘金https://juejin.cn › 后端\n2019年1月3日 — 为此LevelDB 提供了批处理功能，批处理操作就好比事务，LevelDB 确保这一些列写操作的原子性执行，要么全部生效要么完全不生效。 class WriteBatch { void ...\n\nLevelDB源码解析5. 数据完整性 - 知乎专栏https://zhuanlan.zhihu.com › ...\n2018年3月17日 — WAL LOG文件，或者说binlog。是用来支持事务完整性的。 $ cat 000003.log JS )KeyNameExample ValueExample. 对于WAL ...\n\n第四十章LevelDB与BoltDB - 《Go语言四十二章经》"

func main() {
	//leveldb
	client, err := myleveldb.NewClient("/ago/pzxy/WorkSpace/golang/tedge/core/edgex/cmd/tedge-resource/dbdata/992286/")
	if err != nil {
		log.Fatalln(err)
	}
	client.Demo()
}
