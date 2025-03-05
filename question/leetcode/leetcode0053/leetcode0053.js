

const maxSubArray = (nums) => {
  let max = nums[0];

  for (let i = 1; i < nums.length; i++) {
    if (nums[i - 1] + nums[i] > nums[i]) {
      nums[i] = nums[i - 1] + nums[i];
    }
    if (nums[i] > max) {
      max = nums[i];
    }
  }

  return max
}