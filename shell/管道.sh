# 管道
# 匿名管道 pipe，使用过就没有了
# 不能跨终端 ，也就是不能跨进程
rpm -qa |grep bash

# 命名管道，不是常规文件，里面的数据读完就没有了
mkfifo /tmp/tmpfifo
file /tmp/tmpfifo
# print： /tmp/tmpfifo：fifo (named pipe)
# 我们往 /tmp/tmpfifo 文件里面写数据，然后再启动一个进程来读数据。
