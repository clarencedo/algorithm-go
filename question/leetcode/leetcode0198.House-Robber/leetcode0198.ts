function rob(nums: number[]): number {
	const n = nums.length;
	const dp: number[] = new Array(n).fill(0);
	dp[0] = nums[0];
	dp[1] = nums[1] > nums[0] ? nums[1] : dp[0];

	for (let i = 2; i < n; i++) {
		dp[i] = Math.max(dp[i - 1], dp[i - 2] + nums[i])
	}

	return dp[n - 1];
};

function robII(nums: number[]): number {
	const n = nums.length;
	if (n === 0) {
		return 0;
	}
	if (n === 1) {
		return nums[0];
	}

	let prePre: number = nums[0], pre: number = Math.max(nums[0], nums[1]);

	for (let i = 2; i < n; i++) {
		let cur = Math.max(pre, prePre + nums[i])
		prePre = pre;
		pre = cur;
	}

	return pre;
};
