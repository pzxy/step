#!/usr/bin/env bash

if [ $(ps -ef | grep -c "ssh") -eq 1 ];
then
echo "true";
fi

#if else-if else 语法格式：
:<<EOF
if condition1
then
    command1
elif condition2
then
    command2
else
    commandN
fi
EOF


#判断两个变量是否相等：
a=10
b=20
if [ $a == $b ]
then
    echo "a 等于 b"
elif [ $a -lt $b ]
then
    echo "a 小于 b"
else
    echo "a 大于 b"
fi


#if else语句经常与test命令结合使用，

num1=$[2*3]
num2=$[1+5]
if test $[num1] -eq $[num2]
then
    echo "两个数字相等"
else
    echo "两个数字不想等"
fi




