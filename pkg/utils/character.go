package utils

import (
	"math/rand"
	"time"
)

const MIN_CHARACTER = 1
const MAX_CHARACTER = 23

func GetCharacter() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return MIN_CHARACTER + r1.Intn(MAX_CHARACTER+1-MIN_CHARACTER)
}
