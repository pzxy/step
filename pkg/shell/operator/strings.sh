#!/usr/bin/env bash

#下表列出了常用的字符串运算符，假定变量 a 为 "abc"，变量 b 为 "efg"：
#
#运算符	说明	举例
#=	检测两个字符串是否相等，相等返回 true。	[ $a = $b ] 返回 false。
# !=	检测两个字符串是否相等，不相等返回 true。	[ $a != $b ] 返回 true。
#-z	检测字符串长度是否为0，为0返回 true。	[ -z $a ] 返回 false。
#-n	检测字符串长度是否为0，不为0返回 true。	[ -n "$a" ] 返回 true。
#$	检测字符串是否为空，不为空返回 true。	[ $a ] 返回 true。

a="abc"
b="efg"

if [ $a = $b ]
then
   echo "$a = $b : a 等于 b"
else
   echo "$a = $b: a 不等于 b"
fi
if [ $a != $b ]
then
   echo "$a != $b : a 不等于 b"
else
   echo "$a != $b: a 等于 b"
fi
if [ -z $a ]
then
   echo "-z $a : 字符串长度为 0"
else
   echo "-z $a : 字符串长度不为 0"
fi
if [ -n "$a" ]
then
   echo "-n $a : 字符串长度不为 0"
else
   echo "-n $a : 字符串长度为 0"
fi
if [ $a ]
then
   echo "$a : 字符串不为空"
else
   echo "$a : 字符串为空"
fi