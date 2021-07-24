package common

const (
	TarPkg = "./tedge/"
)

const (
	EncryptFileName = "encrypted"
	InputFile       = "./config.yaml"
)

/**
上传文件
*/

//const (
//	User     = "目的主机用户"
//	Host     = "目的主机地址"
//	Port     = "目的主机端口"
//	SrcPath  = "本机文件地址"
//	DstPath  = "目的主机地址"
//	Password = "登陆密码"
//	KeyPath  = "私钥路径"
//	LoadType = "登陆类型"
//)

const (
	ButtonUpload = "上传"
	ButtonPing   = "SSH测试"
)

/**
  SSH
*/

const (
	Darwin  = "darwin"
	Linux   = "linux"
	Windows = "windows"
)

const (
	PassWordType = "密码登录"
	KeyType      = "公钥登陆"
)

const (
	LinuxMacShell  = "ifconfig -a | awk '/^e/  &&  $NF ~ /^([0-9a-fA-F]{2})(([/\\s:][0-9a-fA-F]{2}){5})$/ { print $NF;exit 0}'"
	DarwinMacShell = "ifconfig -a | awk '/ether/  &&  $NF ~ /^([0-9a-fA-F]{2})(([\\s:][0-9a-fA-F]{2}){5})$/ { print $NF;exit 0}'"
)

/**
打包
*/
const (
	HwArch        = "硬件架构"
	Os            = "操作系统"
	PackVersion   = "网关版本"
	PackOutputDir = "安装包输出目录"
)

//https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/1.5.0/edgex_v1.5.0_Darwin_x86_64.tar.gz?raw\=true
//curl -L https://github.com/pzxy/tuya-edge-driver-sdk-go/blob/master/bin/tedge/edgex_v1.5.0_Darwin_arm64.tar.gz\?raw\=true -o edgex_v1.5.0_Darwin_arm64.tar.gz
const (
	PackName                    = "tedge-deploy"
	PackTarName                 = PackName + ".tar"
	PackUnixDefaultOutputDir    = "/tmp/"
	PackWindowsDefaultOutputDir = "C:\\"
	DownloadUrlPrefix           = "https://tedge.coding.net/p/tedge/d/tedge/git/raw/master/"
)

const (
	V1_4_0 = "1.4.0"
	V1_5_0 = "1.5.0"
	V1_6_0 = "1.6.0"
	V1_7_0 = "1.7.0"
	V1_8_0 = "1.8.0"
	V1_9_0 = "1.9.0"
	V2_0_0 = "2.0.0"
)
const (
	V1_4 = "1.4"
	V1_5 = "1.5"
	V1_6 = "1.6"
	V1_7 = "1.7"
	V1_8 = "1.8"
	V1_9 = "1.9"
	V2_0 = "2.0"
)

var VersionList = []string{V1_4_0, V1_5_0, V1_6_0, V1_7_0, V1_8_0, V1_9_0, V2_0_0}
var GwVersionList = []string{V1_4, V1_5, V1_6, V1_7, V1_8, V1_9, V2_0}

func GetGwVersion(s string) string {
	var gw string
	switch s {
	case V1_4:
		gw = V1_4_0
	case V1_5:
		gw = V1_5_0
	case V1_6:
		gw = V1_6_0
	case V1_7:
		gw = V1_7_0
	case V1_8:
		gw = V1_8_0
	case V1_9:
		gw = V1_9_0
	case V2_0:
		gw = V2_0_0
	default:

	}
	return gw
}

/**
上传
*/

const (
	UploadUnixDefaultInFilePath   = PackUnixDefaultOutputDir + PackTarName
	UploadWindowDefaultInFilePath = PackWindowsDefaultOutputDir + PackTarName
)

const (
	TuyaDockerAddr   = "registry.shdocker.tuya-inc.top"
	AliyunDockerAddr = "registry-shdocker-registry.cn-shanghai.cr.aliyuncs.com"
)

var DockerAddrList = []string{AliyunDockerAddr, TuyaDockerAddr}

const (
	DockerRegistry = "镜像仓库"
	GwVersion      = "网关版本"
	DeployAll      = "全部部署"
	TedgeUuid      = "网关uuid"
	EdgeHub        = "tedge-edge-hub"
	Resource       = "tedge-resource"
	Agent          = "tedge-agent"
	Expert         = "tedge-expert"
	GateWay        = "tedge-gateway"
	FileBeat       = "tedge-filebeat"
	WebClient      = "web-client"
)
