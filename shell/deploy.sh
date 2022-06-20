#!/bin/bash
if [ -d "$1" ]; then
  echo "目录已经存在"
  exit
fi
mkdir -p  $1
cp -r {template/conf,template/data,template/logs,template/lg.sh} $1

config="deploy-config.txt"
gatewayId=`sed -r '/^gatewayId:/!d; s/.*://' $config`
secKey=`sed -r '/^secKey:/!d; s/.*://' $config`
localKey=`sed -r '/^localKey:/!d; s/.*://' $config`
port=`sed -r '/^port:/!d; s/^port://' $config`
communityCode=`sed -r '/^communityCode:/!d; s/.*://' $config`
digitalId=`sed -r '/^digitalId:/!d; s/.*://' $config`

sed -r -i "s/gatewayId:.*/gatewayId: $gatewayId/g" $1/conf/base.yml
sed -r -i "s/secKey:.*/secKey: $secKey/g" $1/conf/base.yml
sed -r -i "s/localKey:.*/localKey: $localKey/g" $1/conf/base.yml

sed -r -i "s/gatewayId:.*/gatewayId: $gatewayId/g" $1/conf/config.yml
sed -r -i "s/port:.*/port: $port/g" $1/conf/config.yml
sed -r -i "s/communityCode:.*/communityCode: $communityCode/g" $1/conf/config.yml
sed -r -i "s/digitalId:.*/digitalId: $digitalId/g" $1/conf/config.yml

echo "d=\`date +%Y%m%d\`
tail -f /mnt/anjubao/$1/logs/tuya.log.docker_control.$gatewayId.\${d}.log \$1 \$2
" > $1/lg.sh
sed -r -i "s/\"//g" $1/lg.sh

port2=`echo "$port" | sed -r "s/\://g"`
echo "#
  $1:
    image: "registry-shdocker-registry.cn-shanghai.cr.aliyuncs.com/balloon/anjubao:${2:-"v1.0.8"}"
    network_mode: "host"
    restart: always
    #deploy:
    #  restart_policy:
    #    condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "30m"
        max-file: "10"
    volumes:
      - ./$1/conf/:/home/docker/conf/anjubao/
      - ./$1/logs/:/home/docker/logs/
      - ./$1/data/:/home/docker/tuya.db/
    ports:
      - "$port2:$port2"" > docker-$1.yml