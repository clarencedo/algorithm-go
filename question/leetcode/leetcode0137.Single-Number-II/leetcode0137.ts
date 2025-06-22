function singleNumber(nums: number[]): number {
	// 状态机 ,状态压缩
	// ones 出现1次的位
	// twos 出现2次的位
	// 当某位出现3次时，ones 和twos都会被清掉
	let ones = 0, twos = 0;
	for (const num of nums) {
		ones = (ones ^ num) & ~twos;
		twos = (twos ^ num) & ~ones;
	}

	return ones;
};
function singleNumberII(nums: number[]): number {
	// 对于每一位（32位正数），我们统计所有数字在该位上出现的次数
	// 出现了3次的位，mod 3 之后就会为0
	// 只有那个出现了1次的数字，在它在某些位mod 3后会剩1
	let res = 0;
	for (let i = 0; i < 32; i++) {
		let bitCount = 0;
		for (const num of nums) {
			//取num的第i位
			bitCount += (num >> i) & 1;
		}

		if (bitCount % 3 !== 0) {
			// js中最高位需要特殊处理（负数）
			if (i === 31) {
				res -= (1 << 31);
			} else {
				res |= (1 << i);
			}
		}
	}

	return res;
}