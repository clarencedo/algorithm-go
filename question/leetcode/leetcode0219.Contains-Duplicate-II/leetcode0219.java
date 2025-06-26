import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

class Solution {
	public boolean containsNearbyDuplicate(int[] nums, int k) {
		Map<Integer, Integer> map = new HashMap<>();// key 是num,val 是index

		for (int i = 0; i < nums.length; i++) {
			Integer preIndex = map.get(nums[i]);
			if (preIndex != null && i - preIndex <= k) {
				return true;
			}
			map.put(nums[i], i);
		}

		return false;

	}

	public boolean containsNearbyDuplicateII(int[] nums, int k) {
		Set<Integer> window = new HashSet<>();// 窗口内没有重复元素

		for (int i = 0; i < nums.length; i++) {
			if (i > k) {
				window.remove(nums[i - k - 1]);// 窗口之外就移除
			}
			if (!window.add(nums[i])) {
				return true;
			}
		}

		return false;

	}
}