function majorityElement(nums: number[]): number[] {
	const n = nums.length / 3;
	const freqMap = new Map<number, number>();

	for (const num of nums) {
		freqMap.set(num, (freqMap.get(num) || 0) + 1);
	}

	return [...freqMap].filter(([key, freq]) => {
		return freq > n;
	}).map(([key]) => key);
};