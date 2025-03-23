function maxArea(height: number[]): number {
	let left = 0, right = height.length - 1;
	let capacity = 0;
	while (left < right) {
		let curr = Math.min(height[left], height[right]) * (right - left);
		if (height[left] < height[right]) {
			left++;
		} else {
			right--;
		}
		capacity = Math.max(capacity, curr)
	}

	return capacity;
};