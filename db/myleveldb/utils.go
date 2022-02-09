package myleveldb

import (
	"bytes"
	"encoding/binary"
)

func int64ToBytes(n int64) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes2Int64(bys []byte) int64 {
	buf := bytes.NewBuffer(bys)
	var data int64
	binary.Read(buf, binary.BigEndian, &data)
	return data
}
