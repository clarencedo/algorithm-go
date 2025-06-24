function bucketSort(arr: number[]): number[] {
	if (arr.length === 0) return [];

	const n = arr.length;
	const buckets: number[][] = Array.from({ length: n }, () => []);

	// 1. 将元素分配到桶中
	for (const num of arr) {
		const index = Math.floor(num * n); // 根据范围将元素映射到桶
		buckets[index].push(num);
	}

	// 2. 对每个桶排序
	for (const bucket of buckets) {
		bucket.sort((a, b) => a - b);
	}

	// 3. 合并所有桶
	return buckets.flat();
}