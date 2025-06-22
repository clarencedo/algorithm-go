function findDisappearedNumbers(nums: number[]): number[] {
	const numSet = new Set(nums);
	const n = nums.length;
	const ans: number[] = [];

	for (let i = 1; i <= n; i++) {
		if (!numSet.has(i)) {
			ans.push(i);
		}
	}

	return ans;
};

// 不需要额外空间，在原数组上标记
function findDisappearedNumbersII(nums: number[]): number[] {
	const ans: number[] = [];

	for (let i = 0; i < nums.length; i++) {
		const index = Math.abs(nums[i]) - 1;
		if (nums[index] > 0) {
			nums[index] = -nums[index];
		}
	}

	for (let i = 0; i < nums.length; i++) {
		if (nums[i] > 0) {
			ans.push(i + 1);
		}

	}

	return ans;
};