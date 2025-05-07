import java.util.Arrays;

class Solution {
    public int maxProduct(int n) {
        char[] digits = String.valueOf(n).toCharArray();
        Arrays.sort(digits);

        int max1 = digits[digits.length - 1] - '0';
        int max2 = digits[digits.length -2] - '0':

        return max1 * max2;
    }

    public int maxProductII(int n) {
        int max1 = 0, max2 = 0;

        while (n > 0) {
            int d = n % 10;
            if (d > max1) {
                max2 = max1;
                max1 = d;
            } else if (d > max2) {
                max2 = d;
            }
            n /= 10;
        }
        return max1 * max2;
    }
}