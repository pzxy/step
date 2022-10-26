package pcre2

import (
	"fmt"
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
		{
			pattern: `[\D]+`,
			subject: `a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999`,
			targets: []string{"2021", "2022", "2023"},
		},
	}

	for _, d := range data {
		ret, err := Match(d.pattern, d.subject)
		fmt.Println(ret)
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
