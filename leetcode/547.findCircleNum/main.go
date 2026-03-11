package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	isVisited := make([]bool, n)
	provinceNum := 0

	for i := 0; i < n; i++ {
		if !isVisited[i] { // 发现新省份
			provinceNum++

			// BFS 遍历整个省份
			queue := []int{i}
			isVisited[i] = true // 入队时标记

			for len(queue) > 0 {
				fromCity := queue[0]
				queue = queue[1:]

				for toCity, isC := range isConnected[fromCity] {
					if isC == 1 && !isVisited[toCity] {
						isVisited[toCity] = true // 入队时标记
						queue = append(queue, toCity)
					}
				}
			}
		}
	}

	return provinceNum
}
