package random

import (
	"math/rand"
	"time"
)

func GenRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
