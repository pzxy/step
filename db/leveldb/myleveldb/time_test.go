package myleveldb

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	var day = time.Second * 60 * 24 * 7
	//var day2 = time.Millisecond * 1000 * 60 * 24 * 7
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Add(day).Unix())
	fmt.Println(time.Now().Add(time.Millisecond * 60).Unix())
	fmt.Println(time.Now().Add(time.Second).Unix())
	fmt.Println(time.Now().Add(time.Second).Unix())
	fmt.Println(Expiration(context.Background(), time.Second*3))
	fmt.Println(Expiration(context.Background(), time.Microsecond*3000))
	fmt.Println(Expiration(context.Background(), time.Nanosecond*3000000))
}

//TestRegularClean 测试清理规则
func TestRegularClean(t *testing.T) {
	//c, err := NewClient("./db/leveldb/data/")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//go c.regularClean()
	////设置5秒，睡10秒以后 再获取，应该已经被清理过了
	//err = c.PutStringExpire("a", "v", time.Second*10)
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}
	//time.Sleep(time.Second * 100)
	//v, err := c.getValue("a")
	//if err != leveldb.ErrNotFound {
	//	log.Fatalln("TestRegularClean fail ")
	//}
	//fmt.Println(v)
	//Convey("TestRegularClean", t, func() {
	//	Convey("a", func() {
	//		So(leveldb.ErrNotFound, ShouldEqual, err.Error())
	//	})
	//})
	src := []byte{102, 55, 72, 51, 50, 108, 98, 89, 74, 88, 117, 49, 82, 89, 55, 104}
	toString := hex.EncodeToString(src)
	fmt.Println(toString)
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix() + 2592000)
}
