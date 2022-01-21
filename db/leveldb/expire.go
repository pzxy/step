package leveldb

import (
	"context"
	"fmt"
	"github.com/lunny/log"
	"strings"
	"time"
)

var expirePrefix = "ex:"

func genExpire(key string) []byte {
	return []byte(expirePrefix + key)
}

func genOriginByExpire(key string) ([]byte, error) {
	if strings.HasPrefix(key, expirePrefix) {
		return []byte(key[len(expirePrefix):]), nil
	}
	return nil, fmt.Errorf("prefix not found ")
}

func expireP(duration int64) bool {
	return time.Now().Unix() > duration
}

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func Expiration(ctx context.Context, expiration time.Duration) int64 {
	if expiration > 0 {
		if usePrecise(expiration) {
			//"px",
			fmt.Println("px")
			return formatMs(ctx, expiration)
		} else {
			//"ex",
			fmt.Println("ex")
			return formatSec(ctx, expiration)
		}
	}
	return -1
}

//func Expiration2(ctx context.Context, expiration time.Duration) int64 {
//	time.Now().Unix() + int64(expiration)
//}

func formatMs(ctx context.Context, dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		log.Debugf("specified duration is %s, but minimal supported value is %s - truncating to 1ms", dur, time.Millisecond)
		return 1
	}
	return int64(dur / time.Millisecond)
}

func formatSec(ctx context.Context, dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		log.Debugf(
			"specified duration is %s, but minimal supported value is %s - truncating to 1s",
			dur, time.Second,
		)
		return 1
	}
	return int64(dur / time.Second)
}
