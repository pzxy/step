#!/usr/bin/env bash


echo "is normal"
echo "\"转义字符\""

#echo variable
#read 命令从标准输入中读取一行,并把输入行的每个字段的值指定给 shell 变量
read name
echo "$name is a girl"

# echo newline
echo -e "hello \nworld"

#don't wrap 不换行
echo -e "hello! \c"
echo "world"

#The result is directed to the file (>>  is append data)
echo "The result is directed to the file" > ./file_test

#原样输出不转义(single quote)
echo '$name\"'

#displays command execution results
echo `date`

