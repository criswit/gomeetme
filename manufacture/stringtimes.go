package manufacture

import (
	"fmt"
	"math/rand"
	"time"
)

func makeTimeSlice() {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<5; i++ {
		v[i] = rand.Intn(100)
	}
	fmt.Println(v) // [0 28 27 62 63]
}