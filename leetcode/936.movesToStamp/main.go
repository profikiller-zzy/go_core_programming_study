package main

import "fmt"

func movesToStamp(stamp string, target string) []int {
	curStr := []byte(target)
	m, n := len(stamp), len(curStr)
	windowMatches := make([]int, len(target)-len(stamp)+1)
	for start := 0; start <= len(target)-len(stamp); start++ {
		count := 0
		for i := 0; i < m; i++ {
			if curStr[i+start] == stamp[i] {
				count++
			}
		}
		windowMatches[start] = count
	}
	var markedCount int // 标记现在已经变成？的字符数量
	var result []int
	for markedCount < n {
		found := false
		for start := 0; start <= n-m; start++ {
			// 检查窗口是否有效
			if windowMatches[start] == m {
				windowMatches[start] = -1 // 标记该窗口为已经倒推过，下一次则排除这个窗口
				result = append([]int{start}, result...)
				// 应用stamp，将匹配的字符改为?
				for index := start; index < start+m; index++ {
					if curStr[index] != '?' {
						// 更新受影响的窗口
						left := max(0, index-m+1)
						for k := left; k <= min(index, n-m); k++ {
							if windowMatches[k] == -1 {
								continue
							}
							// k为窗口索引，index-k是字符偏移量
							if curStr[k+(index-k)] != stamp[index-k] {
								windowMatches[k]++
							}
						}
						markedCount++
						curStr[index] = '?'
					}
				}
				found = true
				break
			}
		}
		if !found {
			return []int{}
		}
	}
	return result
}

func main() {
	stamp := "aye"
	target := "eyeye"
	fmt.Println(movesToStamp(stamp, target))
}
