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
	OutputDir      = "加密文件输出目录"
)

const (
	ButtonEncrypt   = "加密"
	Opt             = false
	InputFile       = "/Users/pzxy/WorkSpace/golang/step/gui/tuya-deploy/cryptool/config.yaml"
	OutputFile      = "/Users/pzxy/WorkSpace/golang/step/gui/tuya-deploy/cryptool/encrypted"
	EncryptFileName = "encrypted"
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

//https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/1.5.0/edgex_v1.5.0_Darwin_x86_64.tar.gz?raw\=true
//curl -L https://github.com/pzxy/tuya-edge-driver-sdk-go/blob/master/bin/tedge/edgex_v1.5.0_Darwin_arm64.tar.gz\?raw\=true -o edgex_v1.5.0_Darwin_arm64.tar.gz
const (
	ButtonUpload       = "上传"
	DownloadUrlPrefix  = "https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/"
	DownloadFilePrefix = "/Users/pzxy/WorkSpace/golang/step/gui/tuya-deploy/cryptool/"
)

/**
  SSH
*/

const (
	Darwin  = "darwin"
	Linux   = "linux"
	Windows = "windows"
)

type LoginType string

const (
	PassWord LoginType = "password"
	Key      LoginType = "key"
)

const (
	LinuxMacShell  = "ifconfig -a | awk '/^e/  &&  $NF ~ /^([0-9a-fA-F]{2})(([/\\s:][0-9a-fA-F]{2}){5})$/ { print $NF;exit 0}'"
	DarwinMacShell = "ifconfig -a | awk '/ether/  &&  $NF ~ /^([0-9a-fA-F]{2})(([\\s:][0-9a-fA-F]{2}){5})$/ { print $NF;exit 0}'"
)

/**
打包
*/
const (
	HwArch      = "硬件架构"
	Os          = "操作系统"
	PackVersion = "网关版本"
)
