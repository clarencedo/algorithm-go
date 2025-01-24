package huawei

func dp(heights []int) int {
	//最长递增子序列, longest increase subsequence
	dpLIS := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dpLIS[i] = 1
		for j := 0; j < i; j++ {
			if heights[j] < heights[i] {
				dpLIS[i] = max(dpLIS[i], dpLIS[j]+1)
			}
		}
	}

	//最长递减子序列, longest decrease subsequence
	dpLDS := make([]int, len(heights))
	for i := len(heights); i >= 0; i-- {
		dpLDS[i] = 1
		for j := len(heights); j > i; j-- {
			if heights[j] < heights[i] {
				dpLDS[i] = max(dpLDS[i], dpLDS[j]+1)
			}
		}
	}

	maxLen := 0
	for i := 0; i < len(heights); i++ {
		maxLen = max(maxLen, dpLIS[i]+dpLDS[i]-1)
	}

	if maxLen == 0 {
		return 0
	}

	return len(heights) - maxLen

}