package main

type shape interface {
	rule() bool
	notify() bool
}

type alert struct {
	handers []shape
}

func (a alert) check() {
	for _, v := range a.handers {
		if v.rule() {
			v.notify()
		}
	}
}

func (a alert) register(s shape) {
	a.handers = append(a.handers, s)
}

func main() {

}
