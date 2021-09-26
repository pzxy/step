#!/usr/bin/env bash

:<<!
while循环用于不断执行一系列命令，也用于从输入文件中读取数据；命令通常为测试条件。其格式为：

while condition
do
    command
done
!

#测试条件是：如果int小于等于5，那么条件返回真。
# int从0开始，每次循环处理时，int加1。
# 运行上述脚本，返回数字1到5，然后终止。

int=1
while(($int<=5))
do
 echo $int
 let "int++"
done


#while循环可用于读取键盘信息。
# 输入信息被设置为变量FILM，按<Ctrl-D>结束循环。

echo  '按下 <CTRL-D> 退出'
echo -n '输入你最喜欢的网站名字：'
while read FILE
do
 echo "是的！ $FILE 网站很好"
 echo -n '输入你最喜欢的网站名字：'
done

:<<!


无限循环
无限循环语法格式：

while :
do
    command
done
或者

while true
do
    command
done
或者

for (( ; ; ))
!






