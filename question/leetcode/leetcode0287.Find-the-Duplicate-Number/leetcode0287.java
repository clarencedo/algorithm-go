import java.util.HashSet;
import java.util.Set;

class Solution {
	public int findDuplicate(int[] nums) {
		Integer slow = nums[0];
		Integer fast = nums[0];

		do {
			slow = nums[slow];
			fast = nums[nums[fast]];
		} while (slow != fast);

		slow = nums[0];
		while (slow != fast) {
			slow = nums[slow];
			fast = nums[fast];
		}

		return slow;
	}

	public int findDuplicateII(int[] nums) {
		Set<Integer> seen = new HashSet<>();
		for (int num : nums) {
			if (seen.contains(num)) {
				return num;
			}
			seen.add(num);
		}
		return -1;
	}
}