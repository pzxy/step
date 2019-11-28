## container



- head 堆

````
Heap 包为任何实现 heap.Interface 的类型提供堆操作。
堆是具有属性的树，每个节点是其子树中的最小值节点。

树中的最小元素是参数为0的根。

堆是实现优先级队列的常用方式。要构建一个优先级队列，
请将具有（负）优先级的 Heap 接口作为 Less 方法的排序来执行，因此推送会添加项目，而Pop 将从队列中移除最高优先级的项目。
这些例子包括这样的实现；文件example_pq_test.go 具有完整的源代码。

使用堆的话
首先需要init 元数据
然后元数据类型要实现5个方法，分别是 len less  swap pop push

另外的fix方法可以修改优先级
````

- list 链表

````
双向链表
````

- ring 环
````
Ring是圆形列表或环的元素。
Rings没有开始或结束；
指向任何环形元素的指针用作整个环的参考。
空环表示为零环指针。
一个Ring的零值是一个无零值的单元素环。
````