package pcre2

import (
	"fmt"
	"net"
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
			pattern: `\d{4}([^0-9^\s]{3,11})[^\s]`,
			subject: `a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999`,
			targets: []string{"2482y(Y0", "y(Y"},
		},
	}

	for _, d := range data {
		matchIdx := 0
		ret, err := Match(d.pattern, d.subject, &matchIdx)
		t.Logf("data:%+v,ret:%v \t", d, ret)
		if err != nil {
			t.Error(err)
			continue
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
		t.Log(matchIdx)
	}
}

func TestMatchAndSendMsg2US(t *testing.T) {
	pattern := `\d{4}[^0-9^\s]{3,11}\S`
	subject := `a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999`
	matchStartIdx := 0
	ret, err := Match(pattern, subject, &matchStartIdx)
	if err != nil {
		t.Error(err)
		return
	}
	if err = sendMsg2UDPServer(ret[1], 3000); err != nil {
		t.Error(err)
		return
	}
}

func sendMsg2UDPServer(msg string, port int) error {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: port,
	})
	if err != nil {
		return fmt.Errorf("connect fail，err: %s", err)
	}
	defer socket.Close()
	sendData := []byte(msg)
	_, err = socket.Write(sendData) // send msg
	if err != nil {
		return fmt.Errorf("send message fail，err: %s", err)
	}
	return nil
}
