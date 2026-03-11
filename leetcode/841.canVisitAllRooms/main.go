package main

func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{0} // 最开始只能进入0号房间
	canEnter := make([]bool, len(rooms))
	canEnter[0] = true
	for len(queue) > 0 {
		curRoom := queue[0]
		queue = queue[1:]
		keys := rooms[curRoom]
		for _, key := range keys {
			if !canEnter[key] {
				canEnter[key] = true
				queue = append(queue, key)
			}
		}
	}
	for _, v := range canEnter {
		if !v {
			return false
		}
	}
	return true
}
