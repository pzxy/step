1. [配置文件官网下载地址](http://download.redis.io/redis-stable/redis.conf)


2. 修改redis配置文件

```bash
#注释掉这部分，这是限制redis只能本地访问
bind 127.0.0.1
#默认yes，开启保护模式，限制为本地访问
protected-mode no
#默认no，改为yes意为以守护进程方式启动，可后台运行，除非kill进程，改为yes会使配置文件方式启动redis失败
daemonize no
#数据库个数（可选），我修改了这个只是查看是否生效。。
databases 16 
#输入本地redis数据库存放文件夹（可选）
dir ./
#redis持久化（可选）
appendonly yes
```

3. 启动命令

```bash
docker pull redis

docker run -p 6379:6379  --privileged=true  --name redis  -v /usr/local/docker/redis.conf:/ago/pzxy/WorkSpace/redis/redis.conf  -v /usr/local/docker/data:/data  -d docker.io/redis:latest redis-server /ago/pzxy/WorkSpace/redis/redis.conf  --appendonly yes

```
