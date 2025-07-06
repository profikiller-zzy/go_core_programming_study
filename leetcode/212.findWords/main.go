package main

type trieNode struct {
	children map[byte]*trieNode
	word     string
}

func buildTrie(words []string) *trieNode {
	root := &trieNode{children: make(map[byte]*trieNode)}
	for _, word := range words {
		cur := root
		for index := 0; index < len(word); index++ {
			ch := word[index]
			if cur.children[ch] == nil {
				cur.children[ch] = &trieNode{
					children: make(map[byte]*trieNode),
				}
			}
			cur = cur.children[ch]
		}
		cur.word = word
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	trie := buildTrie(words)
	m, n := len(board), len(board[0])
	var res []string
	// isVisited 可以通过修改board来实现，节省空间
	var dfs func(row int, col int, node *trieNode)
	dfs = func(row int, col int, node *trieNode) {
		ch := board[row][col]
		child := node.children[ch]
		if child == nil {
			return
		}
		if child.word != "" {
			res = append(res, child.word)
			child.word = "" // 去重
		}
		board[row][col] = '#'
		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
				dfs(newRow, newCol, child)
			}
		}
		board[row][col] = ch // 回溯
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			dfs(row, col, trie)
		}
	}
	return res
}

// findWords1 超时了，每次dfs都是完整的dfs
func findWords1(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	isVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		isVisited[i] = make([]bool, n)
	}

	var dfs func(word string, index int, row int, col int) bool
	dfs = func(word string, index int, row int, col int) bool {
		isVisited[row][col] = true

		if index == len(word)-1 {
			isVisited[row][col] = false // ✔ 回溯
			return true
		}

		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
				!isVisited[newRow][newCol] && board[newRow][newCol] == word[index+1] {
				if dfs(word, index+1, newRow, newCol) {
					isVisited[row][col] = false // ✔ 回溯
					return true
				}
			}
		}

		isVisited[row][col] = false // 回溯（必须保证无论成功/失败都执行）
		return false
	}

	type pair struct {
		x, y int
	}
	mapCh2Pair := map[byte][]pair{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if mapCh2Pair[board[row][col]] == nil {
				mapCh2Pair[board[row][col]] = []pair{pair{x: row, y: col}}
			} else {
				mapCh2Pair[board[row][col]] = append(mapCh2Pair[board[row][col]], pair{x: row, y: col})
			}
		}
	}
	var res []string
	for _, word := range words {
		ch := word[0]
		if pairs, ok := mapCh2Pair[ch]; ok {
			for _, p := range pairs {
				if dfs(word, 0, p.x, p.y) {
					res = append(res, word)
					break
				}
			}
		}
	}
	return res
}
