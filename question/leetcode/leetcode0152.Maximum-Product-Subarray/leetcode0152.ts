function maxProduct(nums: number[]): number {
	let n = nums.length;
	if (n === 0) {
		return 0
	}
	let maxVal = nums[0], minVal = nums[0], ans = nums[0]
	for (let i = 1; i < n; i++) {
		let tempMaxVal = maxVal;

		maxVal = Math.max(Math.max(tempMaxVal * nums[i], minVal * nums[i]), nums[i]);
		minVal = Math.min(Math.min(tempMaxVal * nums[i], minVal * nums[i]), nums[i]);
		ans = Math.max(ans, maxVal)
	}

	return ans;
};