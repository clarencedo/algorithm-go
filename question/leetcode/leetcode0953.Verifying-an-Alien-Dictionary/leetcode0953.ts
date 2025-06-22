function isAlienSorted(words: string[], order: string): boolean {
	// 创建字母到其顺序的映射
	const orderMap: Record<string, number> = {};
	for (let i = 0; i < order.length; i++) {
		// 字母 -> 顺序权重
		orderMap[order[i]] = i;
	}

	// 比较两个单词是否按外星语顺序排列
	const compareWords = (word1: string, word2: string): boolean => {
		const len1 = word1.length;
		const len2 = word2.length;
		const maxLen = Math.max(len1, len2);

		for (let i = 0; i < maxLen; i++) {
			const char1 = i < len1 ? word1[i] : undefined;
			const char2 = i < len2 ? word2[i] : undefined;

			// 如果 word1 已经结束，word2 还有字符，则 word1 应该在前面
			if (char1 === undefined) return true;
			// 如果 word2 已经结束，word1 还有字符，则 word2 应该在前面
			if (char2 === undefined) return false;

			const order1 = orderMap[char1];
			const order2 = orderMap[char2];

			if (order1 < order2) return true;
			if (order1 > order2) return false;
			// 如果当前字符相同，继续比较下一个字符
		}

		// 所有字符都相同，返回 true
		return true;
	};

	// 检查所有相邻单词是否按顺序排列
	for (let i = 0; i < words.length - 1; i++) {
		if (!compareWords(words[i], words[i + 1])) {
			return false;
		}
	}

	return true;
}