//Package
package config

import (
	"math/rand"
)
//Generate a new RandomKey
func NewRandomKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		//what else can i do if the random fails ?
		panic(err)
	}
	return key
}
