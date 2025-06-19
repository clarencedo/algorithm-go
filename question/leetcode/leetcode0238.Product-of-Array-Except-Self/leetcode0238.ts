//不能用除法，所以不能用 整个数组乘积 / nums[i]
//所以要把ans[i] 拆成  左边所有数乘积 * 左边所有数乘积
function productExceptSelf(nums: number[]): number[] {
	const n = nums.length;
	const ans: number[] = new Array(n).fill(1);

	//  左侧乘积 前缀积
	for (let i = 1; i < n; i++) {
		ans[i] = ans[i - 1] * nums[i - 1]
	}

	// 左侧乘积 后缀积 临时变量
	let right = 1;
	for (let i = n - 2; i >= 0; i--) {
		right *= nums[i + 1];
		ans[i] *= right;
	}

	return ans;
};