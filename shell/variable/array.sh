#!/usr/bin/env bash

# bash支持一维数组（不支持多维数组），并且没有限定数组的大小
#定义数组
#arr_name=("22",value1)
#单独定义
arr_name[0]="2233"
arr_name[1]=3

#读取数组
valuen=${arr_name[n]}
#使用@符号可以获取数组中的所有元素
echo ${arr_name[@]}

#获取长度
length1=${#arr_name[@]}#2
length2=${#arr_name[*]}#2
lengthn=${#arr_name[1]}#1
echo $length1 $length2 $lengthn