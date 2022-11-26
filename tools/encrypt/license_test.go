package encrypt

import (
	"fmt"
	"testing"
)

func TestLicense(t *testing.T) {
	plaintext := `{
    "appId": "cn.fdm.offlicensedemo",
    "issuedTime": 1595951714,
    "notBefore": 1538671712,
    "notAfter": 1640966400,
    "customerInfo": "亚马逊公司"
}`
	license, hash, err := GenLicense(plaintext)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(license)
	fmt.Println()
	fmt.Println(hash)
	fmt.Println()

	//[82 253 252 7 33 130 101 79 22 63 95 15 154 98 29 114]
	//[20 246 115 204 234 124 210 113 17 132 73 68 32 149 205 64 218 57 22 89 108 202 109 28 63 125 81 204 249 237 9 254]
}
