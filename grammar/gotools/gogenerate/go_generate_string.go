//go:generate stringer -type=TaskRelay,Animal
package gogenerate

type Animal int

const (
	Dog Animal = iota
	Pig
	Cat
	Tiger
	Lion
)

type TaskRelay uint32

const (
	InvalidRelay TaskRelay = iota
	MoveAwayRobotTaskRelay
)
