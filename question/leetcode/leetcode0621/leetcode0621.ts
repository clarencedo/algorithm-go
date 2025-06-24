function leastInterval(tasks: string[], n: number): number {
	const freq: number[] = new Array(26).fill(0);

	for (const task of tasks) {
		freq[task.charCodeAt(0) - 'A'.charCodeAt(0)]++;
	}

	//最大频率
	const maxFreq = Math.max(...freq);

	//有多少个任务拥有最大频率
	const maxCount = freq.filter(f => f === maxFreq).length;

	return Math.max((tasks.length), (maxFreq - 1) * (n + 1) + maxCount);

};