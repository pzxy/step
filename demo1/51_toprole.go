package main

import (
	"fmt"
	"math/rand"
	"time"
)

// uint32_t select = sqrt(all_neighbors.size()) + 1;// 5
// uint32_t size = all_neighbors.size();//16
// for (uint32_t i = 0; i < all_neighbors.size(); ++i) {//16
//
//	    uint32_t idx = RandomUint32() % (size - i);//  xx / 16,15,14,13
//	    if (idx < select) {
//	        random_neighbors.push_back(all_neighbors[i]);
//	        select--;
//	    }
//	}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	all_neighbors := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "1"}
	s := 5
	random_neighbors := make([]string, 0)
	for i := 0; i < len(all_neighbors); i++ {
		idx := int(rand.Uint32()) % (len(all_neighbors) - i)
		if idx < s {
			random_neighbors = append(random_neighbors, all_neighbors[i])
			s--
		}
	}
	fmt.Println(random_neighbors)
}
