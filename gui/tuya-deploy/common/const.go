package common

/**
加密
*/
const (
	Pid            = "产品Id"
	Uuid           = "网关分配的Uuid"
	DeviceID       = "网关Id"
	SecKey         = "SecKey"
	LocalKey       = "LocalKey"
	SubDeviceLimit = "激活设备数量"
	MacAddr        = "目的主机MAC地址"
)

const (
	ButtonEncrypt = "加密"
	Opt           = false
	InputFile     = "/Users/pzxy/WorkSpace/golang/step/gui/tuya-deploy/cryptool/config.yaml"
	OutputFile    = "/Users/pzxy/WorkSpace/golang/step/gui/tuya-deploy/cryptool/encrypted"
)

/**
上传文件
*/

const (
	User = "目的主机用户"
	Host = "目的主机地址"
	Port = "目的主机端口"
	Path = "文件地址"
)
const (
	ButtonUpload = "上传"
)

/**
ssh
*/
type CmdKey int

const (
	Invalid CmdKey = iota
	MacHwAddr
	Kernel
	HwArch
)