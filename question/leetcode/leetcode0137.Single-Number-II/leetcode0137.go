package leetcode

func singleNumber(nums []int) int {
	var res int32 = 0
	for i := 0; i < 32; i++ {
		var count int = 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				count++
			}
		}
		// 第i位上，不是3的倍数->肯定是那个只出现一次的数 “贡献”
		if count%3 != 0 {
			// 第i位是1的数
			// 1 <<3 =8 -> 00001000，即只有第三位是1
			// res |= (1 <<i): 把第i位设成1
			res |= (1 << i)
		}
	}
	return int(int32(res)) // 强制转成有符号结果
}

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}
