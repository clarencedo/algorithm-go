// 维护一个minFreq，只保留共同出现的字符
// 这样最后遍历minFreq就可以输出共同出现的字符
// 以及根据出现的次数 输出重复的字符
function commonChars(words: string[]): string[] {
	if (words.length === 0) {
		return [];
	}
	//以第一个单词位准
	const minFreq: Record<string, number> = {};
	for (const ch of words[0]) {
		minFreq[ch] = (minFreq[ch] || 0) + 1;
	}

	//遍历后续单词，更新最小频率
	for (let i = 1; i < words.length; i++) {
		const currFreq: Record<string, number> = {};
		for (const ch of words[i]) {
			currFreq[ch] = (currFreq[ch] || 0) + 1;
		}

		for (const ch in minFreq) {
			if (currFreq[ch] === undefined) {
				// 当前单词没有该字符，移除
				delete minFreq[ch];
			} else {
				minFreq[ch] = Math.min(minFreq[ch], currFreq[ch]);
			}
		}
	}

	const ans: string[] = [];
	for (const ch in minFreq) {
		for (let i = 0; i < minFreq[ch]; i++) {
			ans.push(ch);
		}
	}

	return ans;
};