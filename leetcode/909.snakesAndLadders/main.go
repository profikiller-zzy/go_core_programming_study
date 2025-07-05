package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	type state struct {
		value int // 当前位置
		step  int // 走到当前位置使用了多少步数
	}
	queue := []state{
		{
			value: 1,
			step:  0,
		},
	}
	isVisited := make([]bool, n*n+1)
	isVisited[1] = true
	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]
		for index := 1; index <= 6; index++ {
			nextValue := curState.value + index
			if nextValue > n*n { // 超出边界
				break
			}
			row, col := getPosition(n, nextValue)
			if board[row][col] != -1 { // 当前位置可以传送
				nextValue = board[row][col]
			}
			if nextValue == n*n {
				return curState.step + 1
			}
			if !isVisited[nextValue] { // 扩展当前状态，将当前格子标记为已检查
				queue = append(queue, state{
					value: nextValue,
					step:  curState.step + 1,
				})
				isVisited[nextValue] = true
			}
		}
	}
	return -1
}

func getPosition(n, value int) (int, int) {
	var col int
	rowFromBottom := (value - 1) / n
	actualRow := n - 1 - rowFromBottom
	posInRow := (value - 1) % n

	if rowFromBottom%2 == 0 {
		col = posInRow
	} else {
		col = n - 1 - posInRow
	}
	return actualRow, col
}
