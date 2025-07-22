package main

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) <= 1 {
		return 0
	}
	state := make([]int, 2*k)
	for index := 0; index < 2*k; index++ {
		if index%2 == 0 {
			state[index] = -prices[0]
		} else {
			state[index] = 0
		}
	}

	for i := 1; i < len(prices); i++ {
		state[0] = max(state[0], -prices[i])
		for j := 1; j < 2*k; j++ {
			if j%2 == 0 {
				state[j] = max(state[j], state[j-1]-prices[i])
			} else {
				state[j] = max(state[j], state[j-1]+prices[i])
			}
		}
	}
	return state[2*k-1]
}
