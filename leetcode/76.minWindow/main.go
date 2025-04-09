package main

import "fmt"

// https://leetcode.cn/problems/minimum-window-substring/description
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	// 先使用一个map存储t中每个字符的出现次数
	tCharMap := make(map[byte]int)
	tCharNumMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCharMap[t[i]]++
		tCharNumMap[t[i]]++
	}

	sCharIndexMap := make(map[byte][]int) // 使用一个map存储s中属于t的字符的位置
	tCharArray := make([]int, 0)          // 使用一个数组存储当前覆盖子串中t的字符的位置
	for index := 0; index < len(s); index++ {
		if _, ok := tCharMap[s[index]]; ok { // 判断s中的字符是否在t中
			sCharIndexMap[s[index]] = append(sCharIndexMap[s[index]], index)
			tCharArray = append(tCharArray, index) // 将s中属于t的字符的位置存储到数组中
		}
	}

	// 核心思想是滑动窗口，首先最小的覆盖子串肯定首尾字符都是t中的字符
	// 那么我们可以从s中找到t的第一个字符的位置，然后从这个位置开始向后遍历
	// 直到找到一个覆盖子串，这个覆盖子串的长度不一定是最小的
	// 之后像滑动窗口一样，将最左边的字符移出，找到下一个覆盖子串
	var right, left, curLen, minLen int
	var minString string
	left = 0
	minLen = len(s)
	for index := 0; index < len(tCharArray); index++ {
		tCharNumMap[s[tCharArray[index]]]--
		if _, ok := tCharMap[s[tCharArray[index]]]; ok {
			tCharMap[s[tCharArray[index]]]--
			if tCharMap[s[tCharArray[index]]] == 0 {
				delete(tCharMap, s[tCharArray[index]])
			}
			if len(tCharMap) == 0 {
				right = index
				break
			}
		}
	}
	if len(tCharMap) != 0 { // 如果没有找到覆盖子串，直接返回空字符串
		return ""
	}
	minLen = tCharArray[right] - tCharArray[left] + 1
	minString = s[tCharArray[left] : tCharArray[right]+1]
	for right < len(tCharArray)-1 { // 不断移动滑动窗口
		if tCharNumMap[s[tCharArray[left]]] < 0 { // 假如左边这个字符的数量大于t中该字符的数量，则右边界不需要向右边移动
			tCharNumMap[s[tCharArray[left]]]++
			left++
			curLen = tCharArray[right] - tCharArray[left] + 1
			if curLen < minLen {
				minLen = curLen
				minString = s[tCharArray[left] : tCharArray[right]+1]
			}
		} else { // 将左边的字符移出之后还需要再右边重新找到这个字符
			curLeftChar := s[tCharArray[left]]
			for right < len(tCharArray)-1 {
				right++
				if s[tCharArray[right]] == curLeftChar { // 找到下一个覆盖子串
					tCharNumMap[s[tCharArray[right]]]--
					tCharNumMap[s[tCharArray[left]]]++ // 将左边的字符移出
					left++
					curLen = tCharArray[right] - tCharArray[left] + 1
					if curLen < minLen {
						minLen = curLen
						minString = s[tCharArray[left] : tCharArray[right]+1]
					}
					break
				}
				tCharNumMap[s[tCharArray[right]]]--
			}
		}
	}

	// 已经找到了右边界为最右边的覆盖子串，继续检查是否有更小的覆盖子串
	for left < len(tCharArray)-1 {
		if tCharNumMap[s[tCharArray[left]]] < 0 {
			tCharNumMap[s[tCharArray[left]]]++
			left++
			curLen = tCharArray[right] - tCharArray[left] + 1
			if curLen < minLen {
				minLen = curLen
				minString = s[tCharArray[left] : tCharArray[right]+1]
			}
		} else {
			break
		}
	}
	return minString
}

func main() {
	s := "bba"
	t := "ab"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}
