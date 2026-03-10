package main

import "fmt"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	left := 0 // left 记录压缩字符串待写入的位置
	curCh := chars[0]
	count := 1

	for right := 1; right < len(chars); right++ { // right 是当前正在遍历的字符位置
		ch := chars[right]
		if ch == curCh {
			count++
		} else {
			chars[left] = curCh
			left++

			if count > 1 {
				countStr := []byte(fmt.Sprintf("%d", count))
				for _, c := range countStr {
					chars[left] = byte(c)
					left++
				}
			}

			curCh = ch
			count = 1
		}
	}

	chars[left] = curCh
	left++
	if count > 1 {
		countStr := []byte(fmt.Sprintf("%d", count))
		for _, c := range countStr {
			chars[left] = byte(c)
			left++
		}
	}

	return left
}
