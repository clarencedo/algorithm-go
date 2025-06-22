function uncommonFromSentences(s1: string, s2: string): string[] {
	const freqMap = new Map<string, number>();
	const ans: string[] = [];

	for (const word of s1.split(' ')) {
		freqMap.set(word, (freqMap.get(word) || 0) + 1)
	}
	const secondMap = new Map<string, number>();
	for (const word of s2.split(' ')) {
		freqMap.set(word, (freqMap.get(word) || 0) + 1)
	}

	//收集在两个句子中只出现一次的词
	for (const [word, freq] of freqMap.entries()) {
		if (freq === 1) {
			ans.push(word);
		}
	}

	return [...freqMap]
		.filter(([_, val]) => val === 1).map(([key]) => key);
	// return ans;
};