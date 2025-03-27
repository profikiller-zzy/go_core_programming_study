package main

//https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	var ans []int
	length := len(s)
	wordCount := len(words)
	wordLength := len(words[0])
	for i := 0; i < wordLength && i+wordCount*wordLength <= length; i++ { // i是滑动窗口的起始位置，i+wordCount*wordLength是滑动窗口的结束位置
		wordMap := make(map[string]int)
		for j := 0; j < wordCount; j++ { // 处理最开始滑动窗口内的单词
			wordMap[s[i+j*wordLength:i+(j+1)*wordLength]]++
		}
		for j := 0; j < wordCount; j++ {
			wordMap[words[j]]--
			if wordMap[words[j]] == 0 {
				delete(wordMap, words[j])
			}
		}

		// 不断将滑动窗口向前移动
		for start := i; start < length-wordCount*wordLength+1; start += wordLength { // wordCount*wordLength是滑动窗口的长度
			// 每次向前滑动一个单词的长度
			if start != i {
				// 每次滑动窗口向前移动一个单词的长度，添加进新的单词++，删除旧的单词--
				newWord := s[start+(wordCount-1)*wordLength : start+wordCount*wordLength]
				wordMap[newWord]++
				if wordMap[newWord] == 0 {
					delete(wordMap, newWord)
				}
				oldWord := s[start-wordLength : start]
				wordMap[oldWord]--
				if wordMap[oldWord] == 0 {
					delete(wordMap, oldWord)
				}
			}
			if len(wordMap) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return ans
}

func main() {

}
