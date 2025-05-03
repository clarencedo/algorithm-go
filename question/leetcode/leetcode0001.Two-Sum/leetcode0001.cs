using System.Collections.Generic;

public class Solution
{
	public int[] TwoSum(int[] nums, int target)
	{
		Dictionary<int, int> table = new Dictionary<int, int>();
		for (int i = 0; i < nums.Length; i++)
		{
			int complement = target - nums[i];
			if (table.ContainsKey(complement))
			{
				return new int[] { table[complement], i };

			}
			table[nums[i]] = i;

		}
		return new int[0];

	}
}
