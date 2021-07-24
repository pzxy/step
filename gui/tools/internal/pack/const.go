package pack

import "example.com/m/v2/common"

/**
打包
*/
const (
	HwArch        = "硬件架构"
	Os            = "操作系统"
	PackOutputDir = "安装包输出目录"
)

//https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/1.5.0/edgex_v1.5.0_Darwin_x86_64.tar.gz?raw\=true
//curl -L https://github.com/pzxy/tuya-edge-driver-sdk-go/blob/master/bin/tedge/edgex_v1.5.0_Darwin_arm64.tar.gz\?raw\=true -o edgex_v1.5.0_Darwin_arm64.tar.gz
const (
	PackTarName = common.PackTarName

	PackDeployFileNamePath = common.TarPkg + "deploy.sh"
	PackReadmeFileNamePath = common.TarPkg + "部署操作方法.md"
)
