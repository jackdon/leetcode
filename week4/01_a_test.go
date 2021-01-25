package week4

import "testing"

func TestCloseStrings(t *testing.T) {
	var (
		word1 = "abc"
		word2 = "bca"
	)
	t.Log(closeStrings(word1, word2))
}
