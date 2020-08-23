package main

import "fmt"

type NodeState int

const (
	InvalidNodeState NodeState = iota
	CommonNodeState
)

func (s NodeState) String() string {
	switch s {
	case CommonNodeState:
		return "CommonNodeState"
	case InvalidNodeState:
		return "InvalidNodeState"
	default:
		return "Invalid type"
	}
	return "Invalid type"
}

type Node struct {
	nodeState NodeState
}

func main() {
	n := new(Node)
	n.nodeState = 1
	// 这里应该是 InvalidNodeState
	fmt.Printf("%+v \n", n.nodeState)
	fmt.Printf("%+v \n ", n)
	fmt.Printf("%v \n", n)
}
