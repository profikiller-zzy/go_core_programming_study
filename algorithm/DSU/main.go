package main

// 并查集(DSU disjoint set union)模板

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 初始时每个元素的parent都是自己
		size[i] = 1
	}
	return &DSU{
		parent, size,
	}
}

// Find 查找到 x 所属集合的根结点
func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // 路径压缩
	}
	return d.parent[x]
}

// Union 合并 a 和 b 所属的集合，返回 true 表示成功合并，返回 false 表示它们原本已经属于同一集合
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

// Same 判断 a、b 是否属于同一集合
func (d *DSU) Same(a, b int) bool {
	return d.Find(a) == d.Find(b)
}

// Size 返回 x 所属集合的元素数量
func (d *DSU) Size(x int) int {
	root := d.Find(x)
	return d.size[root]
}
