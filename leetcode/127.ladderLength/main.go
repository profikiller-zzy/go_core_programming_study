package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	n := len(wordList)
	graph := make([][]bool, n)
	for index := 0; index < n; index++ {
		graph[index] = make([]bool, n)
	}
	type state struct {
		curWordIndex int
		length       int
	}
	var queue []state
	isVisited := make([]bool, n)

	// 构建图
	for index := 0; index < n-1; index++ {
		curWord := wordList[index]
		if curWord == beginWord {
			queue = append(queue, state{curWordIndex: index, length: 1})
			isVisited[index] = true
		}
		if isRelated(beginWord, curWord) {
			queue = append(queue, state{curWordIndex: index, length: 2})
			isVisited[index] = true
		}
		for nextIndex := index + 1; nextIndex < n; nextIndex++ {
			if isRelated(curWord, wordList[nextIndex]) {
				graph[nextIndex][index] = true
				graph[index][nextIndex] = true
			}
		}
	}
	if wordList[n-1] == beginWord {
		queue = append(queue, state{curWordIndex: n - 1, length: 1})
		isVisited[n-1] = true
	}
	if isRelated(beginWord, wordList[n-1]) {
		queue = append(queue, state{curWordIndex: n - 1, length: 2})
		isVisited[n-1] = true
	}

	// bfs
	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]
		if wordList[curState.curWordIndex] == endWord {
			return curState.length
		}
		for index := 0; index < n; index++ {
			if !isVisited[index] && graph[curState.curWordIndex][index] {
				queue = append(queue, state{curWordIndex: index, length: curState.length + 1})
				isVisited[index] = true
			}
		}
	}
	return 0
}

func isRelated(word1, word2 string) bool {
	var dis int
	for index := 0; index < len(word1); index++ {
		if word1[index] != word2[index] {
			dis++
		}
	}
	return dis == 1
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
