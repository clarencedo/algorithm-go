function countSubstrings(s: string): number {
	const n = s.length;
	let count = 0;
	const dp: boolean[][] = Array.from({ length: n }, () => Array(n).fill(false));

	for (let i = n - 1; i >= 0; i--) {
		for (let j = i; j < n; j++) {
			if (s[i] === s[j]) {
				if (j - i <= 1 || dp[i + 1][j - 1]) {
					dp[i][j] = true;
					count++
				}

			}
		}
	}

	return count;
};

//中心扩散法
function countSubstringsII(s: string): number {
	let count = 0;
	for (let center = 0; center < 2 * s.length - 1; center++) {
		let left = Math.floor(center / 2);
		let right = left + (center % 2);

		while (left >= 0 && right < s.length && s[left] == s[right]) {
			count++;
			left--;
			right++;
		}

	}

	return count;
};
