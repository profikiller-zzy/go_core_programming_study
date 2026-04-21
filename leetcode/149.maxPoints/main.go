package main

import (
	"fmt"
	"math"
)

type line struct {
	slope, offset float64
}

func maxPoints1(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	calSlope := func(point1, point2 []int) (float64, float64) {
		x1, y1 := point1[0], point1[1]
		x2, y2 := point2[0], point2[1]
		if x1 == x2 { // 对应函数方程为 x=b 的情况，返回斜率和偏移量
			return math.MaxFloat64, float64(x1)
		}
		return float64(y2-y1) / float64(x2-x1), (float64(x2*y1) - float64(x1*y2)) / float64(x2-x1)
	}
	slopePointMap := make(map[line]map[int]struct{})
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			point1, point2 := points[i], points[j]
			slope, offset := calSlope(point1, point2)
			l := line{slope, offset}
			if _, ok := slopePointMap[l]; !ok {
				slopePointMap[l] = make(map[int]struct{})
			}
			slopePointMap[l][i] = struct{}{}
			slopePointMap[l][j] = struct{}{}
		}
	}
	var maxPointNum int
	for _, p := range slopePointMap {
		if len(p) > maxPointNum {
			maxPointNum = len(p)
		}
	}
	return maxPointNum
}

func maxPoints(points [][]int) (ans int) {
	for i, p := range points { // 假设直线一定经过 points[i]
		x, y := p[0], p[1]
		cnt := map[float64]int{}
		for _, q := range points[i+1:] {
			dx, dy := q[0]-x, q[1]-y
			k := math.MaxFloat64
			if dx != 0 {
				k = float64(dy) / float64(dx)
			}
			cnt[k]++
			ans = max(ans, cnt[k]) // 这里没有算上 (x,y) 这个点，最后再加一
		}
	}
	return ans + 1
}

func main() {
	points := [][]int{{0, 0}, {1, 1}, {1, 0}, {2, 1}}
	fmt.Println(maxPoints(points))
}
