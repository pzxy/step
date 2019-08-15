package main

import (
	"fmt"
	"time"
)

func main() {
	testTime2()
}
func testTime2() {
	//声明时间变量
	var t time.Time
	tt := time.Unix(0, time.Now().Add(20*time.Second).UnixNano())
	ttt := tt.Format("2006-01-02")
	fmt.Println(ttt)
	fmt.Println(t)
}
func testTime() {
	//声明时间变量
	var t time.Time
	fmt.Println(t)
	//获取当前时间
	tNow := time.Now()
	fmt.Println(tNow)
	//时间戳
	tNano := time.Unix(0, tNow.UnixNano()) //tNow.UnixNano()时间转为纳秒   time.Unix(0,tNow.UnixNano())纳秒转为时间
	fmt.Println(tNano, tNow.UnixNano())
	//自己指定时间
	tDate := time.Date(1970, 1, 1, 23, 59, 59, 0, time.Local) //time.Local时区
	fmt.Println(tDate)
	//年月日
	fmt.Println(tDate.Year())
	fmt.Println(tDate.Month())
	fmt.Println(tDate.Day())

	fmt.Println(tDate.Date())
	y, m, d := tDate.Date()
	fmt.Println(y, m, d)
	//时分秒
	fmt.Println(tDate.Hour())
	fmt.Println(tDate.Minute())
	fmt.Println(tDate.Second())
	fmt.Println(tDate.Nanosecond()) //纳秒

	fmt.Println(tDate.Clock())
	//秒差
	fmt.Println(tDate.Unix()) //秒差 距离1970年1月1日
	fmt.Println(tDate.UnixNano())

	//时间格式转换
	//1 time转string
	f := tDate.Format("2006-01-02 15:04:05") //格式随便写 时间必须是这个时间 1 2 3 4 5
	fmt.Println(f)
	//2 string 转time
	str := "2019-01-25 12:01:25"
	t, err := time.Parse("2006-01-02 15:04:05", str) //两者格式应完全一样
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}
