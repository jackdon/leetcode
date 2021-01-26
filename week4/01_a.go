package week4

func closeStrings(word1 string, word2 string) bool {
	// operation 1: swap any two existing characters
	a1, a2 := []rune(word1), []rune(word2)
	lenA1, lenA2 := len(a1), len(a2)
	if lenA1 <= 1 || lenA2 <= 1 {
		return false
	}
	var operation1 = func(a []rune, sizeOfA, n int) bool {
		return false
	}
	if operation1(a1, lenA1, 0) {
		return true
	}
	var operation2 = func(a []rune) bool {
		return false
	}
	return operation2(a1)
}
