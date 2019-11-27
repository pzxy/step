#!/usr/bin/env bash

#使用变量
your_name="freedom"
echo $your_name
echo ${your_name}

# readonly variable
my_name="change"
readonly my_name
readonly dulei="ssss"
echo $dulei

#delete variable
my_lang="China"
#删除后的变量不能使用，unset不能删除readonly变量
unset my_lang
echo $my_lang


