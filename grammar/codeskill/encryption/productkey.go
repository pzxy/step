package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"os/exec"
	"strings"
)

func GetMachineCode() string {
	out, err := exec.Command("/bin/bash", "-c" /*const not safe*/, "dmidecode -s system-uuid && dmidecode -t processor |grep ID | head -1 | awk -F 'ID:' '{print $2}' && dmidecode -s baseboard-serial-number && dmesg | grep eth0 | head -1 | awk '{print $NF}'").Output()
	if err != nil {
		return "-1"
	}

	ids := strings.Split(string(out), "\n")
	ids = append(ids, "pss")
	//n := len(ids)
	idsstr := strings.Join(ids, "|/|")
	//asd := CaesarEncrypt(idsstr,n)
	product := Encrypt([]byte(idsstr), []byte{19, 95, 8, 27})
	return GetMd5String4pss(string(product), true, true)
}

//生成32位md5字串
func GetMd5String4pss(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}

func CaesarEncrypt(ghajsdsdahj string, movenum int) string {
	length := len(ghajsdsdahj)
	n := length >> 2
	if length&3 == 0 {
		n++
		movenum = int(n)
	}
	var asd = make([]byte, length)
	for k, v := range ghajsdsdahj {
		asd[k] = byte(v) + byte(movenum)
	}
	return GetMd5String4pss(string(asd), true, true)
}
