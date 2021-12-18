package tools

import (
	"math/rand"
	"time"
)

const SlugLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

func GenRandSlug(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = SlugLetters[rand.Intn(len(SlugLetters))]
	}
	return string(b)
}
