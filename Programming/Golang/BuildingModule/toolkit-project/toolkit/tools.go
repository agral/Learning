package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

type Tools struct {
}

func (t *Tools) RandomString(length int) string {
	s, runeSource := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		prime, _ := rand.Prime(rand.Reader, len(runeSource))
		x, mod := prime.Uint64(), uint64(len(runeSource))
		s[i] = runeSource[x%mod]
	}

	return string(s)
}
