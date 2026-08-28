package main

import "fmt"

type state struct {
	step int
	word string
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	wordSet := make(map[string]bool)
	for index := 0; index < n; index++ {
		wordSet[wordList[index]] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	letters := []byte("abcdefghijklmnopqrstuvwxyz")

	queue := []state{{1, beginWord}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		bytes := []byte(cur.word)
		for index := 0; index < len(bytes); index++ {
			curCh := bytes[index]
			for _, ch := range letters {
				if ch == curCh {
					continue
				}
				bytes[index] = ch
				next := string(bytes)
				if wordSet[next] {
					if next == endWord {
						return cur.step + 1
					}
					queue = append(queue, state{cur.step + 1, next})
					delete(wordSet, next)
				}
			}
			bytes[index] = curCh
		}
	}
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
