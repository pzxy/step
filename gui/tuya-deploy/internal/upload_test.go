package internal

import (
	"step/grammar/codeskill/log"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	//edgex_v1.5.0_Darwin_x86_64.tar.gz
	//https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/1.5.0/edgex_v1.5.0_Darwin_x86_64.tar.gz
	err := downloadFile("edgex_v1.5.0_Darwin_x86_64.tar.gz", "1.5.0")
	if err != nil {
		log.ErrLog(err)
	}
}
