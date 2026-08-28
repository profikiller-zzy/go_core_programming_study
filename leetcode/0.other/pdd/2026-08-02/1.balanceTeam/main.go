package main

/*
【拼多多】2026-08-02 第一题：平衡队伍

题目描述：
某体育俱乐部的 n 名队员排成一列，每名队员的类型用字符串中的字符表示：
'A' 或 'B'。

教练需要选出一个连续区间组成队伍。若该区间内 A 类队员数与 B 类队员数相等，
则称该队伍为“平衡队伍”。

请找出平衡队伍的最大人数；若不存在平衡队伍，输出 0。

输入：
第一行：整数 n（1 <= n <= 2 * 10^5）。
第二行：长度为 n 的字符串 s，仅包含字符 'A' 和 'B'。

输出：
一个整数，表示平衡队伍的最大人数；若不存在，输出 0。

示例：
输入：
4
ABAB
输出：
4

输入：
3
AAA
输出：
0

输入：
5
AAABB
输出：
4

说明：
第三个示例中，子串 "AABB" 含有 2 个 'A' 和 2 个 'B'，
长度为 4，是最长的平衡队伍。
*/

func maxBalancedTeam(s string) int {
	firstPos := map[int]int{0: -1}
	// 遇见A +1，B -1，prefix记录当前的数值，firstPos记录当前数值第一次出现的下标
	var prefix, result int
	for index, ch := range s {
		if ch == 'A' {
			prefix++
		} else {
			prefix--
		}
		if pos, ok := firstPos[prefix]; ok {
			result = max(result, index-pos)
		} else {
			firstPos[prefix] = index
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
