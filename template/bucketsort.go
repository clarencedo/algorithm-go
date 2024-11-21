package template

import "sort"

func bucketSort(arr []int) []int {
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

	//桶的数量
	numBuckets := len(arr)
	buckets := make([][]int, numBuckets)

	for _, val := range arr {
		//计算桶索引
		bucketIndex := (val - minVal) * numBuckets / (maxVal - minVal)
		buckets[bucketIndex] = append(buckets[bucketIndex], val)
	}

	//对每个桶内部排序
	for i := range buckets {
		sort.Ints(buckets[i])
	}

	//合并所有桶结果
	sortedArr := []int{}
	for _, bucket := range buckets {
		sortedArr = append(sortedArr, bucket...)
	}

	return sortedArr
}