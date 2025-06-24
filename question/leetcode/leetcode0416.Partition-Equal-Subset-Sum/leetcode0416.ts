function canPartition(nums: number[]): boolean {
	let vals: number = 0;
	for (const num of nums) {
		vals += num;
	}

	if (vals % 2 !== 0) {
		return false;
	}

	const target = vals >> 1;
	const dp: boolean[] = Array(target + 1).fill(false);
	dp[0] = true;

	for (const num of nums) {
		for (let i = target; i >= num; i--) {
			//如果原来能组成i，或者现在能通过 i -  num + num 组成i，都算成功
			dp[i] = dp[i] || dp[i - num];
			//写成dp[i] = dp[i - num],会覆盖掉之前其他num得到的true
		}

		if (dp[target]) {
			return true
		}
	}

	return dp[target];
};