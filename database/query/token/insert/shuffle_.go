package insert

import (
	"math/rand"
)

func shuffle(s string) string {
	srune := []rune(s)
	rand.Shuffle(len(srune), func(i, j int) {
		srune[i], srune[j] = srune[j], srune[i]
	})

	return string(srune)
}
