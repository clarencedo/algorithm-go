/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

var twoSum = function (nums, target) {
	let map = new Map();
	for (let i = 0; i < nums.length; i++) {
		let index = map.get(target - nums[i]);
		if (index !== undefined) {
			return [i, index]
		}
		map.set(nums[i], i)
	}
	return [];
};

var twoSum = function (nums, target) {
	//题目要求返回原数组的索引
	const arr = nums.map((val, idx) => ({ val, idx })); // 记录原值和索引
	arr.sort((a, b) => a.val - b.val); // 按值排序j
	arr.sort()

	let left = 0;
	let right = arr.length - 1;

	while (left < right) {
		const sum = arr[left].val + arr[right].val;
		if (sum > target) {
			right--;
		} else if (sum < target) {
			left++;
		} else {
			return [arr[left].idx, arr[right].idx]; // 返回原始索引
		}
	}
	return [];
};