package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var txContexts = []string{CREATE_TOKEN, MINT_TOKEN, BURN_TOKEN, TRANSFER_TOKEN}

// init runs every time a package is used
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name of len 6 for account
func RandomOwner() string {
	return RandomString(6)
}

// RandomTxContext generates a random currency code for account
func RandomTxContext() string {
	n := len(txContexts)
	return txContexts[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
