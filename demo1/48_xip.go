package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.FormatUint(17726169228580487168, 16)
	fmt.Println(c)
	fmt.Printf("%b\n", uint64(17726169228580487168))
	// 17726169228580487168
	// f60000ff02000000
	// domain type network_type ver network_id 			  zone_id 	 cluster_id group_id node_id
	// 1 		11 10110 		000 000000000000011111111 0000001 	 0000000  	00000000 0000000000

	fmt.Println()
	c2 := strconv.FormatUint(144115188075855897, 16)
	fmt.Println(c2)
	fmt.Printf("%b\n", uint64(144115188075855897))
	// 200000000000019
	// 1000000000000000000000000000000000000000000000000000011001
	b2 := 144115188075855897 & 0x3FFFFFFFFFFFFF
	fmt.Println(b2)
}
