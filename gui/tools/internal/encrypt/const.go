package encrypt

import "example.com/m/v2/common"

const (
	Region   = "区域"
	Env      = "环境"
	GwStatus = "网关Status"
	Saas     = "是否是Saas版本"

	Pid              = "网关Pid"
	Uuid             = "网关分配的Uuid"
	DeviceID         = "网关Id"
	SecKey           = "SecKey"
	LocalKey         = "LocalKey"
	SubDeviceTotal   = "SubDeviceTotal"
	SubDeviceLimit   = "可激活设备数"
	MacAddr          = "目的主机MAC地址"
	DefaultOutputDir = "加密文件默认输出目录"
	OutputDir        = "加密文件输出目录"
)

const (
	EncryptFileName = common.EncryptFileName
	ButtonEncrypt   = "加密"
	OutputFile      = common.TarPkg + EncryptFileName
)

//region AY: m1-ay.iot-dns.com      中国
//region AY: m1-we.iot-dns.com      西欧
//region AZ: m1-az.iot-dns.com     美国
//region EU: m1-eu.iot-dns.com     欧洲
//region IN: m1-in.iot-dns.com        印度
//region A1: m1-a1.iot-dns.com       联通
//region V1: m1-v1.iot-dns.com       虚拟地区（联通音箱设备）
//region UE: m1-ue.iot-dns.com     美东
//region DX:m1-DX.iot-dns.com     四川电信
//region A2: m1-A2.iot-dns.com     国美
//region A1: m1-a3.iot-dns.com        某私有云
const (
	China          = "中国AY"
	WesternEurope  = "西欧WE"
	US             = "美国AZ"
	Europe         = "欧洲EU"
	India          = "印度IN"
	Unicom         = "联通A1"
	VirtualArea    = "虚拟地区(联通音箱设备)V1"
	USEast         = "美东UE"
	SichuanTelecom = "四川电信DX"
	GuoMei         = "国美A2"
	PrivateCloud1  = "某私有云A1"
)

var RegionList = []string{China, WesternEurope, US, Europe, India, Unicom, VirtualArea, USEast, SichuanTelecom, GuoMei, PrivateCloud1}

func GetRegion(value string) string {
	switch value {
	case China:
		return "AY"
	case WesternEurope:
		return "WE"
	case US:
		return "AZ"
	case Europe:
		return "EU"
	case India:
		return "IN"
	case Unicom:
		return "A1"
	case VirtualArea:
		return "V1"
	case USEast:
		return "UE"
	case SichuanTelecom:
		return "DX"
	case GuoMei:
		return "A2"
	case PrivateCloud1:
		return "A1"
	default:
		return ""
	}
}

//da_0|daily 日常。pr_0|pre 预发。pro|prod 线上。

const (
	Daily = "日常daily"
	Pre   = "预发pre"
	Prod  = "线上pro"
)

var EevList = []string{Daily, Pre, Prod}

func GetEnv(value string) string {
	switch value {
	case Daily:
		return "daily"
	case Pre:
		return "pre"
	case Prod:
		return "pro"
	default:
		return ""
	}
}
