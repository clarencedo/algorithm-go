package template

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr

	}

	maxVal, minVal := arr[0], arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}

	//创建计数数组
	count := make([]int, maxVal-minVal+1)

	//统计每个元素出现的次数
	for _, val := range arr {
		count[val-minVal]++
	}

	//构建输出数组
	output := make([]int, 0, len(arr))
	for i, c := range count {
		for c > 0 {
			output = append(output, i+minVal)
			c--
		}
	}

	return output
}