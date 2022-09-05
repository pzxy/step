package encrypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	// rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := GenRsaKey()
	// fmt.Println(string(prvKey))
	// fmt.Println(string(pubKey))
	fmt.Println(string(prvKey))
	fmt.Println()
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	fmt.Println([]byte(data))
	signData, _ := RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println(signData)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")

}
