class Solution {
    public int maxArea(int[] height) {
        Integer left = 0, right = height.length - 1;
        Integer capacity = 0;
        while (left < right) {
            var cur = Math.min(height[left], height[right]) * (right - left);
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
            if (cur > capacity) {
                capacity = cur;
            }
        }

        return capacity;
    }

}
