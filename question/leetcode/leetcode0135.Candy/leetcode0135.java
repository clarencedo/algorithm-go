import java.util.Arrays;

class Solution {
	// 贪心
	// 从左边到右遍历，如果当前比左边高，则当前比左边多1
	// 从右边到左遍历，如果当前比右边高，则当前比右边多1
	public int candy(int[] ratings) {
		int n = ratings.length;
		int[] candies = new int[n];
		Arrays.fill(candies, 1);

		for (int i = 1; i < n; i++) {
			if (ratings[i] > ratings[i - 1]) {
				candies[i] = candies[i - 1] + 1;
			}
		}
		for (int i = n - 2; i >= 0; i--) {
			if (ratings[i] > ratings[i + 1]) {
				candies[i] = Math.max(candies[i], candies[i + 1] + 1);
			}
		}

		int sum = 0;
		for (int candy : candies) {
			sum += candy;
		}

		return sum;
	}
}