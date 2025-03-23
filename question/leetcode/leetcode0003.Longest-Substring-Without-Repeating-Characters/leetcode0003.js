const lengthOfLongestSubstring = (s) => {
	let window = new Map()
	let left = 0, right = 0
	let ans = 0

	while (right < s.length) {
		let rightChar = s[right]
		window[rightChar]++
		window.set(rightChar, (window.get(rightChar) || 0) + 1)
		right++

		while (window.get(rightChar) > 1) {
			let leftChar = s[left]
			window.set(leftChar, window.get(leftChar) - 1)
			left++
		}
		ans = Math.max(ans, right - left)
	}
	return ans
}