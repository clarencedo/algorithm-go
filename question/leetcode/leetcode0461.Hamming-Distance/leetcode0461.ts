function hammingDistance(x: number, y: number): number {
	const xor = x ^ y;
	let distance = 0;
	let n = xor;

	while (n !== 0) {
		distance += n & 1; // 检查最低位是否为1,是就distance+1
		n = n >>> 1;        // 无符号右移1位
	}

	return distance;
}