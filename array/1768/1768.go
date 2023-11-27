package main

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i := 0
	for i < len(word1) && i < len(word2) {
		result += word1[i:i+1] + word2[i:i+1]
		i++
	}
	if i != len(word1) {
		result += word1[i:len(word1)]
	}
	if i != len(word2) {
		result += word2[i:len(word2)]
	}
	return result
}
