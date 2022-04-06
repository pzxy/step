package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("GMT")

	// this should give you time in location
	t := time.Now().In(location)
	t.Year()
	t.Unix()
	date := t.AddDate(0, 0, 30).Unix()
	d := date - t.Unix()
	fmt.Println(d)
	s := "asdasdasdasd"
	sum256 := sha256.Sum256([]byte(s))
	fmt.Println(hex.EncodeToString(sum256[:]))
	fmt.Println(hex.EncodeToString(sum256[:])[:32])

}
