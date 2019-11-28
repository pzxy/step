#!/usr/bin/env bash

:<<EOF
Shell case语句为多选择语句。可以用case语句匹配一个值与一个模式，如果匹配成功，执行相匹配的命令。case语句格式如下：
case 值 in
模式1)
    command1
    command2
    ...
    commandN
    ;;
模式2）
    command1
    command2
    ...
    commandN
    ;;
esac
case工作方式如上所示。取值后面必须为单词in，每一模式必须以右括号结束。
取值可以为变量或常数。匹配发现取值符合某一模式后，其间所有命令开始执行直至 ;;。
取值将检测匹配的每一个模式。一旦模式匹配，则执行完匹配模式相应命令后不再继续其他模式。
如果无一匹配模式，使用星号 * 捕获该值，再执行后面的命令。
EOF



echo '输入 1 到 4 的数字： '
echo '您输入的数字为：'
read NUM
case $NUM in
1) echo '您选择了 1'
;;
2) echo '您选择了 2'
;;
*) echo '相当于default'
;;
esac



