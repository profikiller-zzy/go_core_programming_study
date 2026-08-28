package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
题目描述：

多多在一家电商平台做物流调度。平台在 N 个城市之间建立了 M 条有向运输线路。每条线路从城市 u 到城市 v，需要支付邮费 w 元。
一位顾客在城市 1 下单，商品需要从城市 1 运送到城市 N。多多手中有一张免邮券，可以免除任意一条有向运输线路的邮费，即将所选择线路的邮费变为 0。
请计算从城市 1 到城市 N 的最小总邮费。如果即使使用免邮券也无法到达城市 N，则输出 -1。

输入描述：

第一行输入两个整数 N 和 M，分别表示城市数量和有向运输线路数量。

2 <= N <= 100000
0 <= M <= 200000

接下来 M 行，每行输入三个整数 u、v、w，表示存在一条从城市 u
到城市 v 的有向运输线路，通过该线路需要支付邮费 w。

1 <= u, v <= N
1 <= w <= 10000


输出描述：

输出一个整数，表示从城市 1 到城市 N 的最小总邮费。

如果无法从城市 1 到达城市 N，则输出 -1。


示例：

输入：
4 4
1 2 2
1 3 5
2 4 3
3 4 1

输出：
1

示例解释：

选择路径：

1 -> 3 -> 4

原本需要支付的邮费为：

5 + 1 = 6

对费用为 5 的线路 1 -> 3 使用免邮券后，总邮费变为：

0 + 1 = 1
*/

type edge struct {
	from, to int
	price    int64
}

type toNode struct {
	node  int
	price int64
}

type state struct {
	node int
	used int // 从城市1到当前城市最优路径中是否使用了优惠券
	cost int64
}

type priorityQueue struct {
	data []state
}

func (q *priorityQueue) Len() int      { return len(q.data) }
func (q *priorityQueue) Swap(i, j int) { q.data[i], q.data[j] = q.data[j], q.data[i] }

func (q *priorityQueue) UpAdjust(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if q.data[parentIndex].cost > q.data[index].cost {
			q.Swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

func (q *priorityQueue) DownAdjust(index int) {
	n := q.Len()
	for {
		leftChild, rightChild := index*2+1, index*2+2
		bestIndex := index
		if leftChild < n && q.data[leftChild].cost < q.data[bestIndex].cost {
			bestIndex = leftChild
		}
		if rightChild < n && q.data[rightChild].cost < q.data[bestIndex].cost {
			bestIndex = rightChild
		}
		if bestIndex == index {
			break
		}
		q.Swap(bestIndex, index)
		index = bestIndex
	}
}

func (q *priorityQueue) Push(value state) {
	q.data = append(q.data, value)
	q.UpAdjust(q.Len() - 1)
}

func (q *priorityQueue) Pop() state {
	top := q.data[0]
	q.Swap(0, q.Len()-1)
	q.data = q.data[:q.Len()-1]
	q.DownAdjust(0)
	return top
}

func minimumPostage(n, m int, edges []edge) int64 {
	graph := make([][]toNode, n+1)
	for _, e := range edges {
		graph[e.from] = append(graph[e.from], toNode{e.to, e.price})
	}
	// 这里dist[index][0]存储从城市1到index没有使用优惠券的最小费用，dist[index][1]是使用了优惠券的
	dist := make([][2]int64, n+1)
	for index := 1; index <= n; index++ {
		dist[index][0], dist[index][1] = math.MaxInt64, math.MaxInt64
	}
	dist[1][0] = 0
	pq := &priorityQueue{
		data: make([]state, 0),
	}
	pq.Push(state{node: 1, cost: 0, used: 0})
	for pq.Len() > 0 {
		cur := pq.Pop()
		if cur.cost != dist[cur.node][cur.used] {
			// curState是小根堆的堆顶，这里如果不相等，代表这里的curState是过时的
			continue
		}

		for _, next := range graph[cur.node] {
			// case1 :正常支付当前边的费用
			newCost := cur.cost + next.price
			if newCost < dist[next.node][cur.used] {
				dist[next.node][cur.used] = newCost
				pq.Push(state{
					node: next.node,
					cost: newCost,
					used: cur.used,
				})
			}

			// case2: 使用免邮券免费掉当前边的费用
			if cur.used == 0 && cur.cost < dist[next.node][1] {
				dist[next.node][1] = cur.cost
				pq.Push(state{
					node: next.node,
					cost: cur.cost,
					used: 1,
				})
			}
		}
	}
	result := min(dist[n][0], dist[n][1])
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &edges[i].from, &edges[i].to, &edges[i].price)
	}
	fmt.Fprintln(out, minimumPostage(n, m, edges)) // 1
}
