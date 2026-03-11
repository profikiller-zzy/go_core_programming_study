package main

func predictPartyVictory(senate string) string {
	n := len(senate)
	Rqueue := make([]int, 0)
	Dqueue := make([]int, 0)

	// 初始化队列
	for i, ch := range senate {
		if ch == 'R' {
			Rqueue = append(Rqueue, i)
		} else {
			Dqueue = append(Dqueue, i)
		}
	}

	// 模拟投票过程
	for len(Rqueue) > 0 && len(Dqueue) > 0 {
		// 比较队首，索引小的先行动
		if Rqueue[0] < Dqueue[0] {
			// R 先行动，禁止 D
			Rqueue = append(Rqueue, Rqueue[0]+n) // 重新入队（下一轮）
			Rqueue = Rqueue[1:]                  // 移除当前
			Dqueue = Dqueue[1:]                  // 禁止 D（移除）
		} else {
			// D 先行动，禁止 R
			Dqueue = append(Dqueue, Dqueue[0]+n) // 重新入队（下一轮）
			Dqueue = Dqueue[1:]                  // 移除当前
			Rqueue = Rqueue[1:]                  // 禁止 R（移除）
		}
	}

	if len(Rqueue) > 0 {
		return "Radiant"
	}
	return "Dire"
}
