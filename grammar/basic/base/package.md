## gopath信息
>使用命令行env查看

GOARCH="amd64"//目标处理器架构

GOBIN=""//编译器和连接器的安装位置

GOCACHE="/home/pzxy/.cache/go-build"

GOEXE=""

GOFLAGS=""

GOHOSTARCH="amd64"

GOHOSTOS="linux"

GOOS="linux"//目标操作系统

GOPATH="/home/pzxy/gitPush:/home/pzxy/gitPush/godemo:/home/pzxy/go"//当前工作目录

GOPROXY=""

GORACE=""

GOROOT="/usr/local/go"//当前工作目录

GOTMPDIR=""

GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"

GCCGO="gccgo"

CC="gcc"

CXX="g++"

CGO_ENABLED="1"

GOMOD=""

CGO_CFLAGS="-g -O2"

CGO_CPPFLAGS=""

CGO_CXXFLAGS="-g -O2"

CGO_FFLAGS="-g -O2"

CGO_LDFLAGS="-g -O2"

PKG_CONFIG="pkg-config"

GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build861973793=/tmp/go-build -gno-record-gcc-switches"

## gopath工程结构
- gosrc 放置源码
- gopkg 中间缓存文件
- gobin 经过go build ,go install或者go get 产生的二进制文件

##  编译源码
 #### 编译
 go install hello //hello为一个包 ,编译完成后默认保存在gopath/bin下
 #### 执行
 在bin目录中执行 ./hello

## 设置gobin目录
export GOBIN=`echo $GOPATH/bin`

当前bin目录下
export GOBIN=`pwd`

## 注意事项

设置GOPAHT时尽量不要设置全局的GOPATH



 
 
