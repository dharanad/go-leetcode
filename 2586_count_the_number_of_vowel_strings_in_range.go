package main

func VowelStringChecker() func(string) bool {
	vowelSet := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	return func(str string) bool {
		_, s := vowelSet[rune(str[0])]
		_, e := vowelSet[rune(str[len(str)-1])]
		return s && e
	}
}
func vowelStrings(words []string, left int, right int) int {
	isVowelString := VowelStringChecker()
	count := 0
	for i := left; i <= right; i++ {
		if isVowelString(words[i]) {
			count++
		}
	}
	return count
}
