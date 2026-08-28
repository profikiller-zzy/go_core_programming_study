package main

/*
pdd 笔试题
描述：
某公司在 n 个城市设有区域分仓，分仓沿一条环形物流干线布局——仓 1 与仓 2 相邻，仓 2 与仓 3 相邻，……，仓 n 与仓 1 相邻（首尾相接构成一个环）。
每个分仓 i 需要补充 ai 件货物（ai 为非负整数）。总部通过"补货班次"完成补货：每个班次中，调度员选择一组分仓进行集中补货，被选中的分仓各补 1 件货物。
由于环形干线上同一班次内相邻两个分仓同时补货会共用同一段传送带从而产生调度冲突，因此每个班次选中的分仓集合必须是环上的独立集——即任何两个被选中的分仓在环上都不能相邻（注意仓 1 与仓 n 也是相邻的）。
每个班次耗时 1 小时。每个分仓 i 必须恰好被补 ai 次。请你计算最少需要多少个班次才能完成全部补货任务。

输入描述
第一行一个正整数 T，表示测试数据组数。
对于每组测试数据：
- 第一行一个正整数 n (1 ≤ n ≤ 2 × 10^5)，表示分仓数量。
- 第二行 n 个非负整数 a1, a2, ..., an (0 ≤ ai ≤ 10^9)，其中 ai 表示分仓 i 需要补货的件数。

输出描述
对于每组测试数据，输出一行一个整数，表示最少需要的班次数。

补充说明
保证所有测试数据的 n 之和不超过 2 × 10^6。

示例 1
输入
1
3
1 1 1
*/

func reStocking(n int, a []int) int {
	if n == 1 {
		return a[0]
	}
	var aSum, neighborSumMax int
	for i := 0; i < n; i++ {
		aSum += a[i]
		neighborSum := a[i] + a[(i+1)%n]
		if neighborSum > neighborSumMax {
			neighborSumMax = neighborSum
		}
	}
	if n%2 == 0 {
		return neighborSumMax
	}
	k := n / 2 // 奇数环中，一个班次最多选择 floor(n/2) 个仓
	return max(neighborSumMax, (aSum+k-1)/k)
}
