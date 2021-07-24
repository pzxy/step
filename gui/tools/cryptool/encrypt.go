package cryptool

import (
	"example.com/m/v2/common"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type EncryptInfo struct {
	Pid            string
	Uuid           string
	DeviceId       string
	SecKey         string
	LocalKey       string
	SubDeviceLimit string
	SubDeviceTotal string
	Env            string
	Region         string
	GwStatus       string
	Saas           string
}

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

func CreateConfigFile(e *EncryptInfo) error {
	if e == nil {
		return fmt.Errorf("EncryptEntry is nil")
	}

	limit, err := strconv.Atoi(e.SubDeviceLimit)
	if err != nil {
		return err
	}
	var gwStatus bool
	var saas bool
	if e.GwStatus == "true" {
		gwStatus = true
	} else {
		gwStatus = false
	}

	if e.Saas == "true" {
		saas = true
	} else {
		saas = false
	}

	y := &Yaml{
		EdgeHub: EdgeHub{
			BaseConfig: BaseConfig{
				Env:    e.Env,
				Region: e.Region,
			},
			GwConfig: GwConfig{
				Pid:      e.Pid,
				Uuid:     e.Uuid,
				GwId:     e.DeviceId,
				SecKey:   e.SecKey,
				LocalKey: e.LocalKey,
				Status:   gwStatus,
			},
			Saas:           saas,
			SubDeviceTotal: 0,
			SubDeviceLimit: limit,
		},
	}

	b, err := yaml.Marshal(y)
	if err != nil {
		return err
	}
	os.Remove(common.InputFile)
	f, err := os.OpenFile(common.InputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
