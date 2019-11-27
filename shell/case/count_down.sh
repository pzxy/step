#!/usr/bin/env bash

# count down 倒计时

for time in `seq 9 -1 0`;do
echo -e -n "\b$time"
sleep 1
done
#什么都不写输出空格
echo