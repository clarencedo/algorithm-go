function reconstructQueue(people: number[][]): number[][] {
	//身高从大到小，同身高按k从小到大
	people.sort((a, b) => {
		if (a[0] !== b[0]) {
			return b[0] - a[0];
		}
		return a[1] - b[1];
	});

	// 再一个一个插入指定位置
	const res: number[][] = [];
	for (const person of people) {
		//插入在第k位
		res.splice(person[1], 0, person);
	}

	return res;
};