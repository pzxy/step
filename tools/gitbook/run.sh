#!/bin/bash

## 拉取gitbook镜像。
#docker pull billryan/gitbook
## 本项目目录下，也就是当前目录下，执行此命令。
#alias gitbook='docker run  --rm -v "$PWD":/gitbook -p 4000:4000 -d billryan/gitbook gitbook'
## 初始化服务
#gitbook init
## 启动服务,启动后,访问localhost:4000
#gitbook serve
## 停止服务
#docker stop 容器名/容器id

set -e
v=$1
if [[ $v == "" ]] ;then
  echo "未传入端口，默认端口为 4000"
  v=4000
fi
n="xxx-gitbook"


checkport() {
  v=$1
  int=1
  while(( $int<=130 ))
  do
    result=$(netstat -na | grep $v | wc -l)
    dResult=$(docker ps -f name=$n | grep $n | wc -l)
      if [ $result -gt 0 -o $dResult -gt 0 ]; then
        echo "等待 $v 端口被释放,$int --> 130s "
        sleep 1
        let "int++"
        continue
      else
        break;
      fi
  done
    result=$(netstat -na | grep $v | wc -l)
    dResult=$(docker ps -f name=$n | grep $n | wc -l)
  if [ $result -gt 0 -o $dResult -gt 0 ]; then
    echo "$v 端口无法被释放"
    exit 1
  fi
}

docker pull billryan/gitbook

GITBOOK="docker run --rm  -v $PWD:/gitbook -p $v:4000 -d billryan/gitbook gitbook "

if [ $(docker ps -f name=$n | grep $n | wc -l)  -gt 0 ] ;then
   docker rm -f $n
fi

checkport $v
$GITBOOK init
checkport $v
$GITBOOK install
checkport $v

GITBOOK="docker run --name $n --rm  -v $PWD:/gitbook -p $v:4000 -d billryan/gitbook gitbook "
$GITBOOK serve

echo "启动成功,稍等几秒后访问 http://localhost:$v"





