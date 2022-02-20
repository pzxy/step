package myleveldb

import (
	"github.com/lunny/log"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"strconv"
	"time"
)

type LevelDB struct {
	db *leveldb.DB
}

func NewClient(dir string) (c *LevelDB, errEdgeX error) {
	client, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		if client, err = leveldb.RecoverFile(dir, nil); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	ldb := &LevelDB{
		db: client,
	}
	return ldb, nil
}

func (l *LevelDB) PutString(key, value string, wo *opt.WriteOptions) error {
	return l.db.Put([]byte(key), []byte(value), wo)
}

func (l *LevelDB) PutStringExpire(key, value string, dur time.Duration) error {
	if err := l.db.Put([]byte(key), []byte(value), nil); err != nil {
		return err
	}
	return l.setExpire(key, dur)
}

func (l *LevelDB) GetString(key string) (string, error) {
	v, err := l.getExpire(key)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (l *LevelDB) Get(key string) ([]byte, error) {
	return l.db.Get([]byte(key), nil)
}

var ss = "Google leveldb简介 - 简书https://www.jianshu.com › ...\n5、支持事务机制。 6、支持数据库快照功能。 7、支持前向和后向遍历数据。 8、默认线程安全（除了WriteBatch等 ...\n\nLevelDB 入门—— 全面了解LevelDB 的功能特性 - 掘金https://juejin.cn › 后端\n2019年1月3日 — 为此LevelDB 提供了批处理功能，批处理操作就好比事务，LevelDB 确保这一些列写操作的原子性执行，要么全部生效要么完全不生效。 class WriteBatch { void ...\n\nLevelDB源码解析5. 数据完整性 - 知乎专栏https://zhuanlan.zhihu.com › ...\n2018年3月17日 — WAL LOG文件，或者说binlog。是用来支持事务完整性的。 $ cat 000003.log JS )KeyNameExample ValueExample. 对于WAL ...\n\n第四十章LevelDB与BoltDB - 《Go语言四十二章经》"

func (l *LevelDB) Demo() {
	for i := 1; i <= 10000000; i++ {
		err := l.PutStringExpire("a"+strconv.Itoa(i), ss, 60*time.Second)
		if err != nil {
			log.Error("插入出错", err)
			continue
		}
		//fmt.Println("插入", i)
	}
	time.Sleep(time.Second * 2)
	for i := 1; i <= 10000000; i++ {
		_, err := l.GetString(strconv.Itoa(i))
		if err != nil {
			log.Error(err)
			continue
		}
		//fmt.Println("查询", v)
	}
}

func (l *LevelDB) Demo2() {
	l.PutStringExpire("asdasd", ss, 60*time.Second)
}

func (l *LevelDB) GetDemo() {
	if v, err := l.db.Get([]byte("aaaa"), nil); err != nil {
		if err == leveldb.ErrNotFound {
			return
		}
		log.Error(err)
	} else {
		log.Debug(v)
	}
}

func (l *LevelDB) setExpire(key string, dur time.Duration) error {
	return l.db.Put(genExpire(key), int64ToBytes(time.Now().Add(dur).Unix()), nil)
}

//判断是否过期
func (l *LevelDB) getExpire(key string) ([]byte, error) {
	//1. 获取时间key
	t, err := l.db.Get(genExpire(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return l.db.Get([]byte(key), nil)
		}
	}
	//是否过期处理
	if expireP(bytes2Int64(t)) {
		l.db.Delete([]byte(key), nil)
		l.db.Delete(genExpire(key), nil)
	}

	return l.db.Get([]byte(key), nil)
}

func (l *LevelDB) getValue(key string) ([]byte, error) {
	return l.db.Get([]byte(key), nil)
}

func (l *LevelDB) cleanExpire(cleanCount int) {
	iter := l.db.NewIterator(util.BytesPrefix([]byte(expirePrefix)), nil)
	for j := 0; j < cleanCount && iter.Next(); j++ {
		v := bytes2Int64(iter.Value())
		k := string(iter.Key())
		log.Println("j:", j)
		log.Println("k:", k)
		if !expireP(v) {
			log.Debugf("not expire")
			continue
		}
		originK, err := genOriginByExpire(k)
		if err != nil {
			break
		}
		log.Infof("origin key:%s", originK)
		l.db.Delete(iter.Key(), nil)
		l.db.Delete(originK, nil)
	}
	iter.Release()
	return
}

//每过一段清理
func (l *LevelDB) regularClean() {
	for {
		l.cleanExpire(500)
		time.Sleep(time.Second * 3)
	}
}
