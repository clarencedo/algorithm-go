package leetcode

//贪心+位运算
//核心思想
//当最低位有一串连续的1， 比如 .....01111
//直接+1，让它变成.....10000, 只用1次操作就抵消了很多步

func minOperations(n int) int {
	res := 0
	for n > 0 {
		//位与运算,如果是是2的幂 则n&(n-1) ==0
		if n&(n-1) == 0 {
			res++
			break
		}
		//找出最低位的1
		// -n是n的二进制取反+1（补码），只留下最低位的1，其他位置都会被清零
		lowbit := n & -n
		//判断是否有连续的两个1
		if (n & (lowbit << 1)) > 0 {
			//有连续的1，比如11，可以加1来跳跃抵消
			n += lowbit
		} else {
			n -= lowbit
		}
		res++
	}
	return res
}
