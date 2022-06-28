package leetcode

/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出:[null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出:[null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数<= 10000
1 <= value <= 10^5
*/
/**
解法使用两个队列，
[]queue 3
[]max   3

[]queue 3 4
[]max   4

[]queue 3 4 9
[]max   9

[]queue 3 4 9 1
[]max   9 1

[]queue 3 4 9 1 8 小于等于就加到后面，大于就替换掉，出列的时候，如果queue[0]==max[0]就删除max[0]
[]max   9 8

[]queue 3 4 9 1 8 5 8
[]max   9 8 5

[]queue 3 4 9 1 8 5 8
[]max   9 8 8
*/
type MaxQueue struct {
}

func Constructor3() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	return 1
}

func (this *MaxQueue) Push_back(value int) {

}

func (this *MaxQueue) Pop_front() int {
	return 1
}
