function findUnsortedSubarray(nums: number[]): number {
	const n = nums.length;
	let max_seen = -Infinity, min_seen = Infinity;
	let left = -1, right = -1;

	//从左往右找右边界
	//最大值追踪
	for (let i = 0; i < n; i++) {
		if (nums[i] >= max_seen) {
			max_seen = nums[i];
		} else {
			right = i;
		}
	}

	//从右边往左边找左边界
	//最小值追踪
	for (let i = n - 1; i >= 0; i--) {
		if (nums[i] <= min_seen) {
			min_seen = nums[i]
		} else {
			left = i
		}
	}

	return right === -1 ? 0 : right - left + 1;
};