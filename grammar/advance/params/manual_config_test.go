package params

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoadManualConfigs(t *testing.T) {
	convey.Convey("test load manual file ", t, func() {
		mconfig := LoadManualConfigs()
		portP := mconfig.MqInfo.Port == 5762
		fmt.Println(portP)
		//convey.So(portP, convey.ShouldBeTrue)
	})
}
