package internal

import (
	"math/rand"
	"time"
)

func RandNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % max
}