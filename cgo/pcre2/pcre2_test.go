package pcre2

import (
	"testing"
)

func TestMatch(t *testing.T) {

	data := []struct {
		pattern string
		subject string
		targets []string
	}{
		{
			pattern: `\d+`,
			subject: `2022 no time to die`,
			targets: []string{"2022"},
		},
		//结果字符串 自身不包含数字和任何类型的空白字符（如空格、回车等等），其长度为 3 至 11 个字符
		//结果字符串 左侧相邻的字符串是4个数字
		//结果字符串 右侧相邻的字符串不为空
		//正则匹配的次数越少越好，尽可能只使用一个正则表达式
		{
			pattern: `\d{4}[^0-9^\s]{3,11}\S`,
			subject: `a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999`,
			targets: []string{"2482y(Y0"},
		},
	}

	for _, d := range data {
		ret, err := Match(d.pattern, d.subject)
		t.Logf("data:%+v,ret:%v \t", d, ret)
		if err != nil {
			t.Error(err)
			return
		}
		if len(ret) != len(d.targets) {
			t.Fatal()
		}
		for idx, v := range ret {
			if d.targets[idx] != v {
				t.Fatal()
				return
			}
		}
	}
}
