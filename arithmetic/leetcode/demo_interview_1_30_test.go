package leetcode

import (
	"errors"
	"fmt"
	"step/grammar/codeskill/log"
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	/**

	LRUCache cache = new LRUCache( 2 )
	cache.put(1, 1)
	cache.put(2, 2)
	cache.get(1) // 返回  1
	cache.put(3, 3) // 该操作会使得密钥 2 作废
	cache.get(2) // 返回 -1 (未找到)
	cache.put(4, 4) // 该操作会使得密钥 1 作废
	cache.get(1)    // 返回 -1 (未找到)
	cache.get(3) // 返回  3
	cache.get(4) // 返回  4
	*/

	cache := NewLRUCache4(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		log.ErrLog(errors.New("should1 value= 1"))
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		log.ErrLog(errors.New("should2 value= -1 "))
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		log.ErrLog(errors.New("should3 value= -1 "))
	}
	if cache.Get(3) != 3 {
		log.ErrLog(errors.New("should4 value= 3"))
	}
	if cache.Get(4) != 4 {
		log.ErrLog(errors.New("should5 value= 4"))
	}
}

func TestRemoveKDigits(t *testing.T) {
	RemoveKdigits("1432219", 3)
}

func TestArray2Tree_Tree2Array(t *testing.T) {
	target := []int{5, 3, 6, -1, -1, 2, 4, -1, -1, -1, -1, 3, 6, -1, -1}
	root := Array2Tree(target)
	ret := serialize2(root)
	for i, _ := range ret {
		//if target[i] != v {
		fmt.Println(i)
		//}
	}
}

func TestString2Int32(t *testing.T) {
	target := int32(^uint32(0) >> 1)
	s := strconv.Itoa(int(target))
	s = "-" + s
	ret := String2Int32(s)

	if ret != 0-target {
		fmt.Println("fail")
	}

}

func TestSerialize(t *testing.T) {
	target := "{5,4,#,3,#,2}"
	root := deserialize(target)
	fmt.Println(root)

}
