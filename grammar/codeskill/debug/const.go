package debug

import "sync"

type Type string

const (
	RcvRawData            Type = "rcvRawData"
	HeartBeatCounter      Type = "heatBeatCounter"
	SendRcvPkgMarshalBody Type = "sendRcvPkgMarshalBody"
)

var types sync.Map

func InitDebugList() {
	types.Store(RcvRawData, &obj{tName: RcvRawData})
	types.Store(HeartBeatCounter, &obj{tName: HeartBeatCounter})
	types.Store(SendRcvPkgMarshalBody, &obj{tName: SendRcvPkgMarshalBody})
}

func (t Type) Write(msg string) {
	if v, ok := types.Load(t); ok {
		v.(*obj).write(msg)
	}
}
