function longestConsecutive(nums: number[]): number {
	let ans = 0;
	const numSet = new Set(nums);

	for (const num of numSet) {
		if (!numSet.has(num - 1)) {
			let currentNum = num;
			let currentLength = 1;
			while (numSet.has(currentNum + 1)) {
				currentNum++;
				currentLength++;
			}

			ans = Math.max(ans, currentLength);

		}
	}

	return ans;
};