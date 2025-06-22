function wordBreak(s: string, wordDict: string[]): boolean {
	const wordSet = new Set(wordDict);
	const dp: boolean[] = new Array(s.length + 1).fill(false);
	dp[0] = true; // dp[i]前i个字符是否可以被字典中的单词完全切分出来

	for (let i = 1; i <= s.length; i++) {
		for (let j = 0; j < i; j++) {
			const word = s.slice(j, i);
			if (dp[j] && wordSet.has(word)) {
				dp[i] = true;
				break;
			}
		}
	}

	return dp[s.length];
};