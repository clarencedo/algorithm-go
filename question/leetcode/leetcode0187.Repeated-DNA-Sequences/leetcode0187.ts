function findRepeatedDnaSequences(s: string): string[] {
	const len = 10;
	const freqMap = new Map<string, number>();
	const ans: string[] = [];

	for (let i = 0; i <= s.length - len; i++) {
		const substr = s.substring(i, i + 10)
		freqMap.set(substr, (freqMap.get(substr) || 0) + 1);

		if (freqMap.get(substr) === 2) {
			ans.push(substr);
		}

	}

	return ans;
};