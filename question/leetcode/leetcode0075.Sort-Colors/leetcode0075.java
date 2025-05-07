class Solution {
    public void sortColors(int[] nums) {
        int left = 0, right = nums.length - 1;

        for (int i = 0; i <= right; i++) {
            if (nums[i] == 0) {
                int temp = nums[i];
                nums[i] = nums[left];
                nums[left] = temp;
                left++;
            }
            if (nums[i] == 2) {
                int tmp = nums[i];
                nums[i] = nums[right];
                nums[right] = tmp;
                right--;
                i--;
            }
        }
    }
}