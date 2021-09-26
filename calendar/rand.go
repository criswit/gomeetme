package calendar

import (
	"math/rand"
	"time"
)

func GenerateRandomHours() []int {
	v := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 1 ; i < 5gg; i++ {
		v[i] = rand.Intn(100)
	}
	return v
}
