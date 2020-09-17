package leetcode

import (
	"strconv"
	"strings"
)

/**
序列化是将一个数据结构或者对象转换为连续的比特位的操作，
进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。
这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示:这与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。
你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明:不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的

*/

type Codec struct {
	l []string
	s string
}

func Constructor37() Codec {
	return Codec{}
}

/**
自己实现
*/
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.s += "#,"
		return this.s
	}
	this.s += strconv.Itoa(root.Val) + ","
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.s
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data != "" {
		this.l = strings.Split(data, ",")
	}
	if this.l[0] == "#" { //之所以这里没有进行越界检查，是因为#刚好和空节点的位置抵消了。只要此处返回，就不会越界。
		this.l = this.l[1:]
		return nil
	}
	v, _ := strconv.Atoi(this.l[0])
	r := &TreeNode{Val: v}
	this.l = this.l[1:]
	r.Left = this.deserialize("")
	r.Right = this.deserialize("")
	return r
}

//方法1 广度优先 1,2,3,#,#,4,5,#,#,#,#
func (this *Codec) serialize1(root *TreeNode) string {
	stack := []*TreeNode{root}
	vals := make([]string, 0)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			stack = append(stack, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}

func (this *Codec) deserialize1(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data, ",")
	idx := 0
	if vals[idx] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++ //因为最后有"#",所以不担心越界。
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}
	}
	return root
}

//方法2：递归 1,2,null,null,3,4,null,null,5,null,null,
//func rserialize(root *TreeNode, str string) string {
//	if root == nil {
//		str += "null,"
//	} else {
//		str += strconv.Itoa(root.Val) + ","
//		str = rserialize(root.Left, str)
//		str = rserialize(root.Right, str)
//	}
//	return str
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	return rserialize(root, "")
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	l := strings.Split(data, ",")
//	for i := 0; i < len(l); i++ {
//		if l[i] != "" {
//			this.l = append(this.l, l[i])
//		}
//	}
//	return this.rdeserialize()
//}
//
//func (this *Codec) rdeserialize() *TreeNode {
//	if this.l[0] == "null" {
//		this.l = this.l[1:]
//		return nil
//	}
//
//	val, _ := strconv.Atoi(this.l[0])
//	root := &TreeNode{Val: val}
//	this.l = this.l[1:]
//	root.Left = this.rdeserialize()
//	root.Right = this.rdeserialize()
//	return root
//}
