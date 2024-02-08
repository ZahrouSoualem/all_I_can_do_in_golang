package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrsntuvwxyz"

func init() {
	// we use this so each time we call the random functions it will generate new values
	rand.Seed(time.Now().UnixNano())
}
func RandomInteger(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k-1)]
		sb.WriteByte(c)
	}
	return sb.String()
}
