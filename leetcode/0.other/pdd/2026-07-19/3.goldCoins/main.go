package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
题目描述：

多多在玩一个横版视角的接金币游戏：控制的角色可以向左或向右移动，
最大移动速度为 1 个单位距离/秒。

地图上总共会出现 N 枚金币，第 i 枚金币会在 Ti 秒时，
准确掉落在坐标 Xi 的位置；如果没有被接住，就会瞬间消失。

为了接住金币，你必须在 Ti 秒时出现在 Xi 坐标。

此外，由于角色接金币需要短暂的硬直时间，如果你接住了金币 A，
想要再去接金币 B，两次接金币的时间间隔必须严格大于两枚金币之间的距离，即：

TB - TA > |XB - XA|

假设游戏开始前，你可以在任意位置等待。
请问整场游戏最多能够接住多少枚金币？

注意：

如果同一时刻、同一坐标有多枚金币同时掉落，最多只可以接住其中一枚。


输入描述：

第一行包含一个整数 N（1 <= N <= 10^5），表示金币的数量。

接下来 N 行，每行包含两个整数 Ti 和 Xi：

1 <= Ti <= 10^9
-10^9 <= Xi <= 10^9

Ti 表示第 i 枚金币出现的时刻，
Xi 表示第 i 枚金币出现的坐标。


输出描述：

输出一个整数，表示最多能够收集到的金币数量。


样例输入：

5
1 2
5 5
2 3
7 4
6 8


样例输出：

3


样例解释：

一种最优的接金币路径为：

金币1 -> 金币2 -> 金币4

总共能够接住 3 枚金币。
*/

type coin struct {
	left, right int64
}

func maxGoldCoins(n int, coins []coin) int {
	sort.Slice(coins, func(i, j int) bool {
		if coins[i].left == coins[j].left {
			return coins[i].right > coins[j].right
		}
		return coins[i].left < coins[j].left
	})
	tails := make([]int64, 0, n)
	for _, c := range coins {
		right := c.right
		if len(tails) == 0 || right > tails[len(tails)-1] {
			tails = append(tails, right)
			continue
		}
		low, high := 0, len(tails)-1
		for low <= high {
			mid := (low + high) / 2
			if tails[mid] < right {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		tails[low] = right
	}
	return len(tails)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	coins := make([]coin, n)
	for i := 0; i < n; i++ {
		var t, x int64
		fmt.Fscan(in, &t, &x)
		coins[i] = coin{
			left:  t - x,
			right: t + x,
		}
	}
	fmt.Fprintln(out, maxGoldCoins(n, coins))
}
