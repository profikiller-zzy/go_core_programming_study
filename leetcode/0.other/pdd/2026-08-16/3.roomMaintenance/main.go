package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
题目描述：

多多负责安排机房设备的维护任务。

现在共有 n 个维护任务。第 i 个任务包含三个参数：

s[i]：任务开始时间。
e[i]：任务结束时间。
v[i]：完成任务后获得的收益。

任务占用的时间区间为：

(s[i], e[i]]

如果一个任务在时间 t 结束，另一个任务恰好在时间 t 开始，
那么这两个任务不冲突。

多多同一时间只能执行一个任务，因此被选择的任务之间不能有
时间重叠。每个任务最多执行一次，多多可以选择任意数量的任务。

请计算多多最多能够获得多少总收益。


输入描述：

第一行输入一个整数 n，表示维护任务的数量。

1 <= n <= 200000

接下来 n 行，每行输入三个整数 s、e、v，分别表示一个任务的
开始时间、结束时间和收益。

0 <= s < e <= 1000000000
1 <= v <= 1000000000

输入中的任务不保证按照开始时间或结束时间排序。

最终答案可能超过 int 范围，在 Go 中需要使用 int64 保存。


输出描述：

输出一个整数，表示多多能够获得的最大总收益。


样例输入：

5
1 3 50
2 5 40
4 6 70
6 7 30
7 9 60


样例输出：

210


样例解释：

可以选择以下四个任务：

任务区间 (1, 3]，收益为 50；
任务区间 (4, 6]，收益为 70；
任务区间 (6, 7]，收益为 30；
任务区间 (7, 9]，收益为 60。

这些任务之间互不冲突，因此总收益为：

50 + 70 + 30 + 60 = 210
*/

type task struct {
	start, end int
	gain       int64
}

func solution(n int, tasks []task) int64 {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].end == tasks[j].end {
			return tasks[i].start < tasks[j].start
		}
		return tasks[i].end < tasks[j].end
	})

	// dp[i] 代表的是[0,i)，不包括任务i的最大收益
	dp := make([]int64, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		cur := tasks[i]
		j := sort.Search(i, func(j int) bool {
			return tasks[j].end > cur.start
		})
		take := dp[j] + cur.gain
		noTake := dp[i]
		dp[i+1] = max(take, noTake)
	}
	return dp[n]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	tasks := make([]task, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &tasks[i].start, &tasks[i].end, &tasks[i].gain)
	}
	fmt.Println(solution(n, tasks))
}
