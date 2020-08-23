package mytime

import (
	"fmt"
	"github.com/pkg/errors"
	"step/utils/helper/conv"
	"step/utils/log"
	"time"
)

// 输出时间精确到秒时，截取字段的下标
var indexOfSecondAccuracy = 19

func GetCurrTimeWithMilliAccuracy() (nowTime time.Time) {
	t := time.Now()

	rs := []rune(t.String())
	second := rs[:indexOfSecondAccuracy]

	milli := t.UnixNano() / 1e6 % 1e3

	strCurrTime := string(second) + "." + conv.FormatInt64(milli)

	var err error
	nowTime, err = time.Parse("2006-01-02 15:04:05", strCurrTime)
	if err != nil {
		log.ErrLog(errors.WithMessage(errors.WithStack(err), fmt.Sprintf("nowTime= %v, strCurrTime= %v, t= %v",
			nowTime, strCurrTime, t.String())))
	}
	return nowTime
}

func GetCurrTime() (nowTime int64) {
	return time.Now().UnixNano()
}
