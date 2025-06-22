function topKFrequent(words: string[], k: number): string[] {
	const countMap = new Map<string, number>();
	const ans: string[] = [];

	for (const word of words) {
		countMap.set(word, (countMap.get(word) || 0) + 1);
	}

	const topK = [...countMap.entries()]
		.sort((a, b) => {
			if (b[1] !== a[1]) {
				return b[1] - a[1]
			}

			//如果频率相同，按照字典序生序排列
			return a[0].localeCompare(b[0]);
		})
		.slice(0, k)
		.map(([key]) => key);

	return topK;
};