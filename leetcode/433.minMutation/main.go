package main

import "fmt"

func minMutation(startGene string, endGene string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}
	type state struct {
		curGeneIndex int
		step         int
	}
	// 定义无向图的邻接矩阵
	n := len(bank)
	graph := make([][]int, n)
	for index := 0; index < n; index++ {
		graph[index] = make([]int, n)
	}

	var queue []state
	isVisited := make([]bool, n)
	// 构建图
	for index := 0; index < n-1; index++ {
		curGene := bank[index]
		if curGene == startGene {
			queue = append(queue, state{index, 0})
			isVisited[index] = true
		}
		if checkMutation(startGene, curGene) {
			queue = append(queue, state{index, 1})
			isVisited[index] = true
		}
		for j := index + 1; j < n; j++ {
			if checkMutation(curGene, bank[j]) {
				graph[index][j] = 1
				graph[j][index] = 1
			}
		}
	}
	if bank[n-1] == startGene {
		queue = append(queue, state{n - 1, 0})
		isVisited[n-1] = true
	}
	if checkMutation(startGene, bank[n-1]) {
		queue = append(queue, state{n - 1, 1})
		isVisited[n-1] = true
	}

	// bfs
	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]
		if bank[curState.curGeneIndex] == endGene {
			return curState.step
		}
		for index := 0; index < n; index++ {
			if !isVisited[index] && graph[curState.curGeneIndex][index] == 1 {
				queue = append(queue, state{index, curState.step + 1})
				isVisited[index] = true
			}
		}
	}
	return -1
}

func checkMutation(startGene string, endGene string) bool {
	var dis int
	n := len(startGene)
	for index := 0; index < n; index++ {
		if startGene[index] != endGene[index] {
			dis++
		}
	}
	return dis == 1
}

func main() {
	startGene := "AACCGGTT"
	endGene := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(startGene, endGene, bank))
}
