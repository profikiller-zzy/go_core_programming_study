package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述：

多多需要处理 T 套告警网络。

每套网络由 N 个节点和 N 条无向边组成，没有自环和重边，
并且任意两个节点之间都可以互相到达。因此，这套网络中恰好存在一个简单环。

节点 i 的初始告警状态为 bi：

- bi = 0，表示告警已关闭；
- bi = 1，表示告警正在响起。

第 i 条边连接节点 ui 和 vi，维护这条边需要花费 ci。

维护一条边时，该边两个端点的告警状态都会翻转：

- 0 变为 1；
- 1 变为 0。

一份维护方案是一个边的集合：

- 集合中的每条边恰好维护一次；
- 未被选中的边不进行维护；
- 边的维护顺序不会影响最终结果；
- 无向边不区分维护方向。

如果执行完方案中的所有维护操作后，所有节点的告警状态都变为 0，
则称该方案为“可行方案”。

一份可行方案的费用等于其中所有被维护边的费用之和。

请计算：

1. 所有可行方案中的最小费用；
2. 达到最小费用的可行方案数量。


输入描述：

第一行包含一个整数 T（1 <= T <= 3），表示测试用例数量。

对于每个测试用例：

第一行包含一个整数 N（3 <= N <= 2 * 10^5），表示节点数量。

第二行包含 N 个整数：

b1, b2, ..., bn

其中 bi ∈ {0, 1}，表示每个节点的初始告警状态。

接下来 N 行，每行包含三个整数：

ui, vi, ci

表示第 i 条无向边连接节点 ui 和 vi，维护费用为 ci。

约束：

1 <= ui, vi <= N
ui != vi
1 <= ci <= 10^9

保证输入图连通，不含自环和重边，并且恰好包含一个简单环。


输出描述：

对于每个测试用例输出一行。

如果不存在可行方案，输出：

-1 0

否则输出两个整数：

最小费用  达到最小费用的可行方案数量


样例输入：

3
5
0 1 1 1 1
1 2 4
2 3 2
3 1 7
3 4 5
4 5 1
4
1 1 1 1
1 2 1
2 3 1
3 4 1
4 1 1
3
1 0 0
1 2 1
2 3 1
3 1 1


样例输出：

3 1
2 2
-1 0


样例解释：

对于第一组测试数据，可以选择第 2、5 条边进行维护，
总费用为：

2 + 1 = 3

另一种可行方案的费用为 12，因此最小费用为 3，
达到最小费用的方案只有 1 种，输出：

3 1
*/

type Edge struct {
	u, v int // 节点u,v
	cost int64
}

// AdjEdge 使用邻接表表示的边
type AdjEdge struct {
	to int
	id int // 边的id
}

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	parent, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA == rootB {
		return false
	}
	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}
	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]
	return true
}

func solve(n int, status []int, edges []Edge) (int64, int) {
	dsu := newDSU(n)
	tree := make([][]AdjEdge, n)
	extraEdgeID := -1

	for id, edge := range edges {
		if dsu.Union(edge.u, edge.v) { // 返回true说明两个节点目前的状态是不连通
			tree[edge.u] = append(tree[edge.u], AdjEdge{to: edge.v, id: id})
			tree[edge.v] = append(tree[edge.v], AdjEdge{to: edge.u, id: id})
		} else {
			extraEdgeID = id
		}
	}

	// 在生成树上建立父子关系
	parent := make([]int, n)     // 节点id -> parent节点ID
	parentEdge := make([]int, n) // 节点id -> edgeID
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -2
		parentEdge[i] = -1
	}

	parent[0] = -1
	order := make([]int, 0, n)
	stack := []int{0} // 从根节点开始访问

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for _, next := range tree[u] {
			if next.to == parent[u] { // 说明这个节点已经检查过了
				continue
			}

			parent[next.to] = u
			parentEdge[next.to] = next.id
			depth[next.to] = depth[u] + 1
			stack = append(stack, next.to)
		}
	}

	// 标记图中唯一环中的边
	onCycle := make([]bool, n)
	onCycle[extraEdgeID] = true
	u, v := edges[extraEdgeID].u, edges[extraEdgeID].v

	for depth[u] > depth[v] {
		onCycle[parentEdge[u]] = true
		u = parent[u]
	}
	for depth[v] > depth[u] {
		onCycle[parentEdge[v]] = true
		v = parent[v]
	}

	for u != v {
		onCycle[parentEdge[u]] = true
		onCycle[parentEdge[v]] = true
		u, v = parent[u], parent[v]
	}

	need := make([]int, n)
	copy(need, status)
	selected := make([]bool, n)
	var baseCost int64

	for i := len(order) - 1; i > 0; i-- {
		u := order[i]
		if need[u] == 1 {
			edgeId := parentEdge[u]
			selected[edgeId] = true
			baseCost += edges[edgeId].cost
			// 选择了这个边之后，相应的父节点状态需要反转
			need[parent[u]] ^= 1
		}
	}
	if need[0] != 0 {
		return -1, 0
	}
	var cycleTotal int64
	var baseCycleCost int64
	for id, edge := range edges {
		if onCycle[id] {
			cycleTotal += edge.cost
			if selected[id] {
				baseCycleCost += edge.cost
			}
		}
	}
	alternativeCost := baseCost + cycleTotal - 2*baseCycleCost
	if alternativeCost == baseCost {
		return alternativeCost, 2
	}
	return min(alternativeCost, baseCost), 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var sampleCount int
	fmt.Fscan(in, &sampleCount)
	for i := 0; i < sampleCount; i++ {
		var n int
		fmt.Fscan(in, &n)
		status := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &status[j])
		}
		edges := make([]Edge, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &edges[j].u, &edges[j].v, &edges[j].cost)
			edges[j].u -= 1
			edges[j].v -= 1
		}
		fmt.Println(solve(n, status, edges))
	}
}
