package main

func maxVowels(s string, k int) int {
	vowelMap := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}
	var count, maxCount int
	for i := 0; i < k; i++ {
		if vowelMap[s[i]] {
			count++
		}
	}
	maxCount = count
	for i := k; i < len(s); i++ {
		if vowelMap[s[i]] {
			count++
		}
		if vowelMap[s[i-k]] {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
