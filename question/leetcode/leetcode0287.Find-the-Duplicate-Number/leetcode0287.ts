function findDuplicate(nums: number[]): number {
	const freqMap = new Map<number, number>();
	for (const num of nums) {
		freqMap.set(num, (freqMap.get(num) || 0) + 1);
	}

	return [...freqMap].sort((a, b) => b[1] - a[1]).slice(0, 1).map(([key]) => key)[0];
};
//不用额外空间
//可以转换为 链表找环的入口的问题
//每个数字都在[1,n]范围内，可以把每个值当成下一个节点的指针，形成一个链表
//因为有重复数，所以一定有环
function findDuplicateII(nums: number[]): number {
	let slow = nums[0];
	let fast = nums[0];

	//第一次相遇
	do {
		slow = nums[slow];
		fast = nums[nums[fast]];
	} while (slow !== fast);

	//寻找环的入口,即重复的那个数
	slow = nums[0];
	while (slow !== fast) {
		slow = nums[slow];
		fast = nums[fast];
	}

	return slow;
};