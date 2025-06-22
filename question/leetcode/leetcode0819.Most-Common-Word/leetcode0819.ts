function mostCommonWord(paragraph: string, banned: string[]): string {
	const words = paragraph.toLowerCase().match(/\b[a-z]+\b/g) || [];
	const countMap: Map<string, number> = new Map();
	for (const word of words) {
		if (!banned.includes(word)) {
			countMap.set(word, (countMap.get(word) || 0) + 1);
		}
	}

	const ans = [...countMap].sort((a, b) => {
		return b[1] - a[1]
	}).slice(0, 1).map(([key]) => key);

	return ans.join();
};