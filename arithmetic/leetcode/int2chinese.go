package leetcode

import (
	"errors"
	"math"
	"step/utils/log"
	"strconv"
)

/**
整数数字变为汉字 10001 一万零一
*/
func int2Chinese(n int) (ret string) {
	//j := len(s) - 1
	s := strconv.Itoa(n)
	j := len(s) - 1
	preChar := math.MaxInt32
	for _, char := range s {
		if int(char) != 48 {
			tmp := translate(int(char), j)
			if preChar == 48 {
				ret += "零"
			}
			ret += tmp
		}

		preChar = int(char)
		j--
	}

	return ret
}

func translate(num, unit int) (s string) {
	switch num {
	case 48:
		s += "零"
		return
	case 49:
		s += "一"
	case 50:
		s += "二"
	case 51:
		s += "三"
	case 52:
		s += "四"
	case 53:
		s += "五"
	case 54:
		s += "六"
	case 55:
		s += "七"
	case 56:
		s += "八"
	case 57:
		s += "九"
	default:
		log.ErrLog(errors.New("num 不能为负数，或者大于10 "))
	}
	switch unit {
	case 0:
		s += ""
	case 1:
		s += "十"
	case 2:
		s += "百"
	case 3:
		s += "千"
	case 4:
		s += "万"
	case 5:
		s += "十万"
	case 6:
		s += "百万"
	case 7:
		s += "千万"
	case 8:
		s += "亿"
	case 9:
		s += "千亿"
	case 10:
		s += "万亿"
	case 11:
		s += "十万亿"
	case 12:
		s += "百万亿"
	case 13:
		s += "兆"
	default:
		log.ErrLog(errors.New("数字太大了"))
	}
	return
}
