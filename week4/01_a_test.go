package week4

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	var (
		word1 = "abc"
		word2 = "bca"
	)
	t.Logf("%s close to %s: %v\n", word1, word2, closeStrings(word1, word2))
}
