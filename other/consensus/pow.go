package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// 每10个块,调整一下出块时间, 最终是控制到,每一分钟出一个块
// 一般情况下应该大一些,比如 2000
const theBlock = 100

var t = (time.Minute * theBlock).Milliseconds()

// 最大难度值
const maxTarget = "00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

func main() {
	difficulty := uint64(16 * 16 * 8)
	target := newTarget(maxTarget, difficulty)
	blockNum := 0
	for {
		start := time.Now().UnixMilli()
		for i := 0; i < theBlock; i++ {
			hash := pow(genStr(10), target)
			blockNum += 1
			fmt.Println(blockNum, time.Now().Format("2006-01-02 15:04:05"), "block:", hash)
		}
		end := time.Now().UnixMilli()
		// 使用旧难度,重新计算难度
		difficulty = newDiff(difficulty, uint64(end-start))
		// 获取新的目标值
		target = newTarget(maxTarget, difficulty)
	}

}

func pow(s string, target string) string {
	src, nonce := s, 0
	for verify(sha(src), target) == false {
		nonce += 1
		src = block(s, nonce)
	}
	return sha(src)
}

func sha(src string) string {
	s := sha256.Sum256([]byte(src))
	ss := sha256.Sum256(s[:])
	return hex.EncodeToString(ss[:])
}

func block(s string, nonce int) string {
	return s + strconv.Itoa(nonce)
}

func verify(src string, target string) bool {
	t1, _ := new(big.Int).SetString(src, 16)
	t2, _ := new(big.Int).SetString(target, 16)
	return t1.Cmp(t2) != 1
}

// newDiff
// 难度计算公式: 旧难度值/新难度值 =( 过去n个区块花费时长 / 预期总时长,比如我们预期1分钟一个块,这里就是n分钟 )
func newDiff(oldDiff uint64, oldTime uint64) uint64 {
	diff := (oldDiff * uint64(t)) / oldTime
	//	 难度不能为 0,如果新难度为 0 说明上次执行的时间很短,我们应该提高难度,难度值越高
	if diff == 0 {
		diff = oldDiff * 16
	}
	fmt.Println(" diff ------------------------> :", diff)
	return diff
}

// newTarget 目标值计算公式 目标值 = 最大目标值 / 难度值
func newTarget(maxTarget string, diff uint64) string {
	t1, _ := new(big.Int).SetString(maxTarget, 16)
	t2 := new(big.Int).SetUint64(diff)
	ret := new(big.Int).Div(t1, t2).Text(16)
	fmt.Println("adjust target: ", len(ret), ret)
	return ret
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
