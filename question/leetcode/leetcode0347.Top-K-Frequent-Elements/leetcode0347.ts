function topKFrequent(nums: number[], k: number): number[] {
	const freqMap = new Map<number, number>();

	for (const num of nums) {
		freqMap.set(num, (freqMap.get(num) || 0) + 1);
	}

	return [...freqMap]
		.sort((a, b) => b[1] - a[1])
		.slice(0, k)
		.map(([key, freq]) => key);
};

//桶排序优化时间复杂度 O(nlogn)
function topKFrequentII(nums: number[], k: number): number[] {
	const freqMap = new Map<number, number>();
	for (const num of nums) {
		freqMap.set(num, (freqMap.get(num) || 0) + 1);
	}

	const bucket: number[][] = Array(nums.length + 1).fill(0).map(() => []);
	for (const [num, freq] of freqMap.entries()) {
		bucket[freq].push(num); // 出现 freq 次的数字放到 bucket[freq]
	}

	const result: number[] = [];
	for (let i = bucket.length - 1; i >= 0 && result.length < k; i--) {
		if (bucket[i].length > 0) {
			result.push(...bucket[i]);
		}
	}

	return result.slice(0, k); // 防止最后一次超过k个元素
}

//用最小堆优化