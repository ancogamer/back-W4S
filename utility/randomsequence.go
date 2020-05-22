package utility

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func Randomsequence() string {
	var sequence []string
	for i := 0; i < 6; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(10))
		sequence = append(sequence, result.String())
	}
	return strings.Join(sequence, "")
}
