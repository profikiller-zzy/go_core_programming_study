package main

func reverseVowels(s string) string {
	vowelSet := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// 收集所有元音的索引
	var vowelIndices []int
	for i := 0; i < len(s); i++ {
		if vowelSet[s[i]] {
			vowelIndices = append(vowelIndices, i)
		}
	}

	// 双指针交换元音
	left, right := 0, len(vowelIndices)-1
	bytes := []byte(s)
	for left < right {
		bytes[vowelIndices[left]], bytes[vowelIndices[right]] =
			bytes[vowelIndices[right]], bytes[vowelIndices[left]]
		left++
		right--
	}

	return string(bytes)
}
