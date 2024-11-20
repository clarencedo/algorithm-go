package leetcode

func missingNumber(arr []int) int {
	n := len(arr)
	diff := (arr[n-1] - arr[0]) / n
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] != diff {
			return arr[i-1] + diff
		}
	}
	return arr[0]
}

// Binary Search
func missingNumberII(arr []int) int {
	left, right := 0, len(arr)-1
	diff := (arr[right] - arr[left]) / len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == arr[0]+mid*diff {
			left = mid + 1
		} else {
			right = mid
		}
		if left == right {
			break
		}
	}
	return arr[0] + left*diff
}