package main

import (
	"fmt"
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	result := make([][]string, 0, len(searchWord))
	prefix := ""

	for i := 0; i < len(searchWord); i++ {
		prefix += string(searchWord[i])

		// 二分查找第一个 >= prefix 的位置
		left := sort.Search(len(products), func(j int) bool { return products[j] >= prefix })
		// 收集最多 3 个匹配的产品
		matches := make([]string, 0, 3)
		for j := left; j < len(products) && len(matches) < 3; j++ {
			if strings.HasPrefix(products[j], prefix) {
				matches = append(matches, products[j])
			} else {
				break // 后面的都不匹配了（已排序）
			}
		}
		result = append(result, matches)
	}

	return result
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	fmt.Println(suggestedProducts(products, "mouse"))
}
