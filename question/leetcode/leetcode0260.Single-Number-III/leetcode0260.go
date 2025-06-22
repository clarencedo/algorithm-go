package leetcode

// 两个数a,b只出现一次，其他数都出现两次，所以异或之后会抵消
// 整个数组异或之后就是 a ^ b
// a != b, 所以a ^ b !=0,它的某一位上是1，说明a,b在那一位上不相同
// 用不同位把原数组分成两类：
// 1. 在该位上是1， 包含a或b中的一个
// 2. 在改为上是0， 包含另一个
// 每类中国呢，除了a , b外的数字都是成对的，异或之后会被抵消
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	// xorSum -> a ^ b
	// -xorSum 是xorSum的补码： 按位取反 +1
	// 一个数的负数是将这个数的二进制表示取反加1
	// 所以 xorSum & -xorSum会保留最右边的那个1，其余全位0
	// 这个操作的作用是：
	// 找出a,b 而今之中某一位不同的位置（即最低的不同位）
	// a = 0110, b= 1010
	// a ^ b = 1100, lsb = 0011 +1 =  0100
	lsb := xorSum & -xorSum
	//利用lsb这一位，将原数组分成两类
	// - 1. 该位是1（包含a或b中的1个）
	// - 2. 该位是0 （包含另一个）
	// 每组中，除了a或b以外，其他都是成对的->异或之后位0
	// 所以最终得到: type1 = a 或b, type2 = b 或a
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
