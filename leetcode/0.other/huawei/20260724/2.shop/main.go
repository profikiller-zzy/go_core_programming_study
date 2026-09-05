package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

/*
题目描述：
每件商品最多买一件。购物原价每满 200 减 20；若所选商品包含至少 3 个不同类别，则改为每满 200 减 30，两种优惠不叠加。
给定预算，求折后价不超过预算时的最大满意度总和。

输入描述：
第一行输入预算（整数），第二行输入商品数 n，随后 n 行输入唯一编号、类别、价格、满意度。
保证预算在 [10, 1000]，1 <= n <= 20，类别在 [1, 10]，价格在 [1, 200]，满意度在 [1, 240]。

输出描述：
输出最大满意度总和。

示例1
输入：
100
3
10001 1 100 90
10002 2 100 100
10003 3 100 110
输出：
110
说明：预算只能购买其中一件。

示例2
输入：
200
4
10001 1 200 200
10002 2 10 10
10003 3 10 10
10004 3 10 10
输出：
230
说明：四件原价 230，三个类别触发减 30，折后正好 200。
*/
type product struct {
	id           int
	kind         int
	price        int
	satisfaction int
}

func solution(budget, n int, products []product) int {
	var dfs func(int, int, int, uint)
	var result int
	dfs = func(index, priceSum, satisfactionSum int, kindMask uint) {
		if index == n {
			// 商品选完了 计算价格
			discount := 20
			kindCount := bits.OnesCount(kindMask)
			if kindCount >= 3 {
				discount = 30
			}
			finalPrice := priceSum - (priceSum/200)*discount
			if finalPrice <= budget && satisfactionSum > result {
				result = satisfactionSum
			}
			return
		}
		// case1 不选择下标为index的商品
		dfs(index+1, priceSum, satisfactionSum, kindMask)

		newKindMask := kindMask | 1<<(products[index].kind-1)
		dfs(index+1, priceSum+products[index].price, satisfactionSum+products[index].satisfaction, newKindMask)
	}
	dfs(0, 0, 0, 0)
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var budget, n int
	fmt.Fscan(in, &budget, &n)
	products := make([]product, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &products[i].id, &products[i].kind, &products[i].price, &products[i].satisfaction)
	}
	// fmt.Println(budget, n, products)
	fmt.Println(solution(budget, n, products))
}
