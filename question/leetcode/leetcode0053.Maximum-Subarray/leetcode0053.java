import java.util.Arrays;

class Solution {
    public int maxSubArray(int[] nums) {
        int max = nums[0];

        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] + nums[i] > nums[i]) {
                nums[i] = nums[i - 1] + nums[i];
            }

            if (nums[i] > max) {
                max = nums[i];
            }
        }

        return max;
    }

    public int maxSubArrayWithPath(int[] nums) {
        int sum_max = nums[0], current_max = nums[0];
        int start = 0, end = 0, temp_start = 0;
        
        for (int i =1; i < nums.length;i++){
            if (current_max + nums[i] > nums[i]){
                current_max += nums[i];
            }else{
                current_max = nums[i];
                temp_start = i;
            }
            
            if (current_max > sum_max){
                sum_max = current_max;
                start = temp_start;
                end = i;
            }
            
        }
        

        int[] path = Arrays.copyOfRange(nums, start, end+1);
        
        return sum_max;
    }
}
