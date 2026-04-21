package main

import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i] < tasks[j] })
	sort.Slice(workers, func(i, j int) bool { return workers[i] < workers[j] })
	n, m := len(tasks), len(workers)
	canSolveKTasks := func(k int) bool {
		// 只需要判断最强壮的k个工人是否能够完成最轻松的k个任务
		index, p := 0, pills
		var validTasks []int
		for _, worker := range workers[m-k:] {
			for index < k && worker+strength >= tasks[index] {
				validTasks = append(validTasks, tasks[index])
				index++
			}
			if len(validTasks) == 0 { // 吃药也完成不了任务
				return false
			}
			// 如果当前工人的力气可以完成最轻松的任务，则让他完成这个任务，直接进入下个循环
			if worker >= validTasks[0] {
				validTasks = validTasks[1:]
				continue
			}
			// 必须吃药
			if p == 0 { // 没药了
				return false
			}
			// 吃药则完成最困难的任务
			p--
			validTasks = validTasks[:len(validTasks)-1]
		}
		return true
	}
	left, right := 0, min(m, n)
	for left <= right {
		mid := (left + right) / 2
		if canSolveKTasks(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}
