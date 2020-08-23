package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
func (node treeNode) print() {
	fmt.Println(node.value)
}

//普通中序列遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

//管道来实现遍历  更序列化 放到了管道中
func (node *treeNode) TraverseFunc(f func(*treeNode)) {
	if node == nil {
		return
	}
	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

func (node *treeNode) traverseWithChan() chan *treeNode {
	out := make(chan *treeNode)
	go func() {
		node.TraverseFunc(func(n *treeNode) {
			out <- n
		})
		close(out)
	}()
	return out
}

func main() {
	var root treeNode
	//结构创建无论地址还是本身都是使用.来访问的
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//创建一个放置节点的切片
	/*nodes := []treeNode{
		{value:3},{},//这里防止多少个值都行
	}
	fmt.Println(nodes)*/
	root.traverse()
	fmt.Println("................")

	//打印管道中的数据
	c := root.traverseWithChan()
	for node := range c {
		fmt.Println(node.value)
	}
}
