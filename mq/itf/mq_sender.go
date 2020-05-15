package itf

type Sender interface {
	Send(msg interface{}) bool
}
