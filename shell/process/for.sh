#!/usr/bin/env bash


:<<EOF
与其他编程语言类似，Shell支持for循环。

for循环一般格式为：

for var in item1 item2 ... itemN
do
    command1
    command2
    ...
    commandN
done
写成一行：

for var in item1 item2 ... itemN; do command1; command2… done;
当变量值在列表里，for循环即执行一次所有命令，使用变量名获取列表中的当前取值。命令可为任何有效的shell命令和语句。in列表可以包含替换、字符串和文件名。

in列表是可选的，如果不用它，for循环使用命令行的位置参数。

EOF

#例如，顺序输出当前列表中的数字：
for loop in 1 2 3 4 5 ; do
    echo "this is value: $loop"
done

#顺序输出当前列表中的字符：

for str in 'this is a string'
do
    echo $str
done





