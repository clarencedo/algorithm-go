package leetcode

func leastBricks(wall [][]int) int {
	//每个边缘位置出现的次数
	edgeCount := make(map[int]int)
	for _, row := range wall {
		//每一行砖块的某个边缘位置的累积宽度
		edge := 0
		//不包括最后一块砖的边缘
		for i := 0; i < len(row)-1; i++ {
			edge += row[i]
			edgeCount[edge]++
		}

	}

	//最多行在某个边缘位置有砖块的边缘
	maxEdge := 0
	for _, count := range edgeCount {
		if count > maxEdge {
			maxEdge = count
		}
	}

	return len(wall) - maxEdge
}
