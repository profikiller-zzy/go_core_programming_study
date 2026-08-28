package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述：

多多在维护一批编号为 1, 2, ..., n 的实例。每个实例只有两种状态：

0 表示使用旧版本；
1 表示使用新版本。

多多在灰度发布平台执行了一次操作：

1. 首先选择一个非空连续区间 [L, R]；
2. 然后选择一个目标状态 V，其中 V 为 0 或 1；
3. 最后将区间 [L, R] 内所有实例的状态都设置为 V。

区间中可以包含原本已经处于状态 V 的实例，但是这次操作必须至少改变
一个实例的状态。

发布前后的实例状态 A 和 B 被完整保留，但操作日志丢失了。
保证至少存在一种合法操作，可以将 A 变为 B。

请判断本次操作能否被唯一确定。

如果合法的三元组 (L, R, V) 恰好只有一个，输出该三元组；
否则输出 -1。

注意：
只要 L、R 或 V 中存在任意一项不同，就视为不同的操作，
即使这些操作最终得到的发布后状态完全相同。


输入描述：

第一行包含一个正整数 T（1 <= T <= 10），表示测试用例的数量。

对于每个测试用例：

第一行包含一个正整数 n，表示实例数量。

第二行包含一个长度为 n 的 01 串 A，表示发布前的状态。

第三行包含一个长度为 n 的 01 串 B，表示发布后的状态。

A 和 B 均为长度恰好为 n 的 01 串。

单个输入文件中，所有测试用例的 n 之和不超过 2 * 10^5。

保证每组数据至少存在一种合法操作。


输出描述：

对于每个测试用例：

如果操作唯一，输出一行三个整数：

L R V

如果不存在唯一操作，输出一行：

-1


样例输入：

6
5
00000
01110
5
01000
01110
6
111111
100001
7
0001000
0111110
1
0
1
5
00010
01110


样例输出：

2 4 1
-1
2 5 0
2 6 1
1 1 1
-1
*/

type operation struct {
	l, r, v int
}

func solution(n int, A, B string) (bool, *operation) {
	left, right := 0, n-1
	for left < n && A[left] == B[left] {
		left++
	}
	if left == n {
		return false, nil
	}
	for right >= 0 && A[right] == B[right] {
		right--
	}

	target := B[left]
	for i := left; i <= right; i++ {
		if B[i] != target {
			return false, nil
		}
	}
	// 尝试看看能不能拓宽边界
	if left > 0 && B[left-1] == target {
		return false, nil
	}
	if right < n-1 && B[right+1] == target {
		return false, nil
	}
	return true, &operation{left + 1, right + 1, int(target - '0')}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for ; testCount > 0; testCount-- {
		var n int
		var A, B string

		fmt.Fscan(in, &n, &A, &B)

		ok, op := solution(n, A, B)
		if !ok {
			fmt.Fprintln(out, -1)
			continue
		}

		fmt.Fprintln(out, op.l, op.r, op.v)
	}
}
