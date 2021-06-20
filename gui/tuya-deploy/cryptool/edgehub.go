package cryptool

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"step/gui/tuya-deploy/const"
	"strconv"
)

type GwConfig struct {
	Pid      string `yaml:"pid"`
	Uuid     string `yaml:"uuid"`
	AuthKey  string `yaml:"authKey"`
	GwId     string `yaml:"gwid"`
	SecKey   string `yaml:"seckey"`
	LocalKey string `yaml:"localkey"`
	Status   bool   `yaml:"status"`
}
type BaseConfig struct {
	Region string `yaml:"region"`
	Env    string `yaml:"env"`
	Atop   string `yaml:"atop"`
	Mqtt   string `yaml:"mqtt"`
}
type EdgeHub struct {
	BaseConfig
	GwConfig
	Saas           bool `yaml:"saas"`
	SubDeviceTotal int  `yaml:"subdevicetotal"`
	SubDeviceLimit int  `yaml:"subdevicelimit"`
}
type Yaml struct {
	EdgeHub
}

func CreateConfigFile(pid, uuid, deviceId, localKey, subDeviceLimit string) error {
	limit, err := strconv.Atoi(subDeviceLimit)
	if err != nil {
		return err
	}
	y := &Yaml{
		EdgeHub: EdgeHub{
			BaseConfig: BaseConfig{
				Env:    "prv",
				Region: "Ay",
			},
			GwConfig: GwConfig{
				Pid:      pid,
				Uuid:     uuid,
				GwId:     deviceId,
				LocalKey: localKey,
			},
			Saas:           true,
			SubDeviceTotal: 0,
			SubDeviceLimit: limit,
		},
	}

	b, err := yaml.Marshal(y)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(_const.InputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
	if err != nil {
		p, _ := os.Getwd()
		fmt.Println(p)
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
