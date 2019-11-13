//go:generate stringer -type Animal
package gogenerate

type Animal int

const (
	Dog Animal = iota
	Pig
	Cat
	Tiger
	Lion
)
