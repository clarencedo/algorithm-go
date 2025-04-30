package leetcode

//	有个数学性质：
//   - x^n = (x^{n/2})^2 当 n 是偶数
//   - ^n = x * (x^{(n-1)/2})^2 当 n 是奇数
//	我们每次都可以把问题规模砍半，像二分一样！
//     这样时间复杂度可以从O(n)降到O(logn)
//     如果n =0,返回1，任何数的0次方都是1
//     如果n<0,先把n变成正数，同时x取倒数
//   - 如果n是偶数，酸楚pow(x,n/2),然后平方
//   - 如果n是奇数，酸楚pow(x,n/2),然后平方再乘x

// 解法一： 迭代
func myPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}
	res := 1.0
	for N > 0 {
		//奇数
		if N%2 == 1 {
			res *= x
		}
		x *= x
		N /= 2
	}

	return res
}

// 解法二： 递归, 最佳解法
func myPowII(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}