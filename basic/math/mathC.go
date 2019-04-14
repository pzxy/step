package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var a, b = 12.3, 45.6
	fmt.Println(a, b)
	fmt.Println(math.Floor(a))  //向下取整
	fmt.Println(math.Ceil(b))   //向上取整
	fmt.Println(math.Abs(a))    //绝对值
	fmt.Println(math.Modf(a))   //取a小数位
	fmt.Println(math.Pow(2, 3)) //2的3次方
	fmt.Println(math.Round(b))  //四舍五入

	/**
	随机数
	*/
	rand.Seed(1)                     //默认种子是1  种子一样每次的随机数就一样
	rand.Seed(time.Now().UnixNano()) //设置当前时间纳秒时间差为种子
	fmt.Println(rand.Intn(10))       //返回0-9的随机数
	fmt.Println(rand.Intn(10))       //返回
	fmt.Println(rand.Intn(10))       //返回

}
