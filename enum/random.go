package enum

import (
	"math/rand"
	"time"
)

func Random[T any](s []T) T {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s))]
}
