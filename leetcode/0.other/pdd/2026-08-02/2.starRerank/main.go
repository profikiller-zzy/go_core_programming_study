package main

import "fmt"

/*
题目描述：

给定 N 条评价的星级 r1, r2, ..., rN，其中每个星级均为 1 到 5。
请将这些星级重新排列，使任意两个相邻元素的值都不相同。

在所有合法的排列中，输出字典序最小的一种；
如果不存在合法排列，则输出 -1。

两个排列 a 和 b 的字典序比较方式如下：
从左到右找到第一个满足 ai != bi 的位置。
若 ai < bi，则排列 a 的字典序小于排列 b。

输入描述：

第一行输入一个整数 N，表示评价数量。
1 <= N <= 10^5

第二行输入 N 个整数 r1, r2, ..., rN，表示每条评价的星级。
1 <= ri <= 5

输出描述：

如果不存在满足条件的排列，输出 -1。

否则输出 N 个整数，整数之间用空格分隔，
表示字典序最小的合法排列。

示例：

输入：
5
1 1 1 2 3

输出：
1 2 1 3 1
*/

func StarRerank(n int, stars []int) []int {
	var starCount [6]int
	for _, star := range stars {
		starCount[star]++
	}
	for _, count := range starCount {
		if count > (n+1)/2 {
			return []int{-1}
		}
	}

	path := make([]int, 0, n)
	var index, prev int
	for ; index < n; index++ {
		found := false
		// 要找到字典序最小的合法序列，需要star从左往右贪心选择
		for star := 1; star <= 5; star++ {
			if star == prev || starCount[star] == 0 {
				continue
			}
			starCount[star]-- // 先贪心选择
			if canFinish(n-len(path), star, starCount) {
				path = append(path, star)
				prev = star
				found = true
				break
			}
			starCount[star]++ // 选了之后后面的非法了，回退
		}
		if !found {
			return []int{-1}
		}
	}
	return path
}

// canFinish 现在剩余n个位置，并且星级评分为starCount的情况下，将第一个位置选为curStar，判断是否存在一个合法的序列
func canFinish(n int, curStar int, starCount [6]int) bool {
	for star := 1; star <= 5; star++ {
		if star == curStar {
			if starCount[star] > n/2 {
				return false
			}
			continue
		}
		if starCount[star] > (n+1)/2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(StarRerank(5, []int{1, 1, 1, 2, 3}))
}
