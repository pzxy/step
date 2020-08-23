package encryption

import (
	"fmt"
	"strings"
	"testing"
)

func TestProductKey(t *testing.T) {
	out := `ba64479c-0000-0000-aa72-6e6172989d83
 00 00 00 00 00 00 00 00
veth5fa88a2
`
	ids := strings.Split(string(out), "\n")
	ids = append(ids, "pss")
	//n := len(ids)
	idsstr := strings.Join(ids, "|/|")
	//asd := CaesarEncrypt(idsstr,n)
	product := Encrypt([]byte(idsstr), []byte{19, 95, 8, 27})
	expect := GetMd5String4pss(string(product), true, true)
	if expect != "01EDAE04021DC4E8" {
		t.Fatal(fmt.Sprintf("expect 01EDAE04021DC4E8 but %v", expect))
	}
}
