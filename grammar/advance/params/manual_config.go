package params

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"step/grammar/codeskill/log"
)

type ManualConfig struct {
	MqInfo rabbmitmq.MqInfo `json:"MqInfo"`
}

func LoadManualConfigs() (manualConfig *ManualConfig) {
	manualConfig = &ManualConfig{}

	bytes, err := ioutil.ReadFile("configfile/configfile.json")
	if err != nil {
		bytes, err = ioutil.ReadFile("../configfile/configfile.json")
		if err != nil {
			bytes, err = ioutil.ReadFile("../../configfile/configfile.json")
		}
	}
	if err != nil {
		log.ErrLog(errors.WithStack(err))
		return
	}
	err = json.Unmarshal(bytes, manualConfig)
	if err != nil {
		return
	}

	return
}
