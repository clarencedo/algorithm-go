function frequencySort(s: string): string {
	const countMap = new Map<string, number>();

	for (const ch of s) {
		countMap.set(ch, (countMap.get(ch) || 0) + 1);
	}

	const sortedEntries = [...countMap.entries()].sort((a, b) => b[1] - a[1]);

	let ans = "";
	for (const [char, count] of sortedEntries) {
		ans += char.repeat(count);
	}

	return ans;
};