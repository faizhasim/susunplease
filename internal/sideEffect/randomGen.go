package sideEffect

import (
	"math/rand"
	"time"
)

type RandomGen interface {
	RandStringBytesMask() string
}

type randomGen struct {
	seed int64
}

const letterBytes = "0123456789abcdef"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func (randomGen *randomGen) RandStringBytesMask() string {
	n := 5
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func NewRandomGen() RandomGen {
	return &randomGen{seed: time.Now().UnixNano()}
}
