package test

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"testing"
)

func TestLeveldb(t *testing.T) {
	// 打开数据库
	db, err := leveldb.OpenFile("/Users/pzxy/WorkSpace/Go/src/step/middleware/leveldb/db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// 增
	db.Put([]byte("1:1"), []byte("1"), nil)
	db.Put([]byte("2"), []byte("2"), nil)
	db.Put([]byte("3"), []byte("3"), nil)
	db.Put([]byte("4"), []byte("4"), nil)
	db.Put([]byte("5"), []byte("5"), nil)
	db.Put([]byte("55"), []byte("55"), nil)

	// //批量操作，增删
	// batch := new(leveldb.Batch)
	// batch.Put([]byte("6"), []byte("6"))
	// batch.Put([]byte("7"), []byte("7"))
	// batch.Put([]byte("8"), []byte("8"))
	// batch.Delete([]byte("4"))
	// err = db.Write(batch, nil)
	// //查
	// v, err := db.Get([]byte("1"), nil)
	// fmt.Println(string(v))
	//
	// //查子集
	// fmt.Println("查子集 2~4,不包括4")
	// sli := &util.Range{Start: []byte("2"), Limit: []byte("56")}
	// iter := db.NewIterator(sli, nil)
	// for iter.Next() {
	//	err := db.Delete(iter.Key(), nil)
	//	if err != nil {
	//		fmt.Errorf("err:%v", err)
	//	}
	//	fmt.Println(string(iter.Key()), string(iter.Value()))
	// }
	// iter.Release()
	//
	// fmt.Println("查子集 2~4,不包括4")
	// sli2 := &util.Range{Start: []byte("16"), Limit: []byte("56")}
	// iter2 := db.NewIterator(sli2, nil)
	// for iter2.Next() {
	//	err := db.Delete(iter2.Key(), nil)
	//	if err != nil {
	//		fmt.Errorf("err:%v", err)
	//	}
	//	fmt.Println(string(iter2.Key()), string(iter2.Value()))
	// }
	// iter2.Release()
	//
	// //从某一点开始查
	// fmt.Println("从2开始查")
	// iter = db.NewIterator(nil, nil)
	// for ok := iter.Seek([]byte("2")); ok; ok = iter.Next() {
	//	fmt.Println(string(iter.Key()), string(iter.Value()))
	// }
	// iter.Release()

	// 根据前缀遍历数据库内容
	fmt.Println("根据前缀1遍历数据库内容  ")
	iter := db.NewIterator(util.BytesPrefix([]byte("1")), nil)
	for iter.Next() {
		fmt.Println(string(iter.Key()), string(iter.Value()))
	}
	// 这个释放是必须的。
	iter.Release()
	err = iter.Error()
}
