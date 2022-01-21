package leveldb

import (
	"context"
	"fmt"
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
