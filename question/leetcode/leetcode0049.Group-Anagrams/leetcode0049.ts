function groupAnagrams(strs: string[]): string[][] {
	const strMap = new Map<string, string[]>();
	for (const str of strs) {
		const ch = str.split('').sort().join('');
		if (!strMap.has(ch)) {
			strMap.set(ch, [])
		}

		strMap.get(ch)!.push(str);
	}

	return Array.from(strMap.values());
};