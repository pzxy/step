#!/usr/bin/env bash

# Single quotes
# 单引号里的任何字符都会原样输出，单引号字符串中的变量是无效的；
echo str='Single quotes string'

#double quotes
#双引号里可以有变量，双引号里可以出现转义字符
your_name='change'
echo -e str="Hello , i know you are \"$your_name\"! \n"

#Concatenated string
#Concatenated by double
d_str="dog"
d_str1="hello，"$d_str"!"
d_str2="hello，"${d_str}"!"
echo $d_str1 $d_str2
#Concatenated by single
echo
s_str="nice"
s_str1='yee，'$s_str'!'
#单引号中不能识别{}
s_str2='yee，${s_str}!'
echo $s_str1 $s_str2

#Get string length
string="i know"
echo ${#string} #6

#Get substring
substring="may the force be with you"
echo ${substring:0:3} # may 包头不包尾

#query substring
string1="runoob is a great site"
echo `expr index "${string1}" io`#4


