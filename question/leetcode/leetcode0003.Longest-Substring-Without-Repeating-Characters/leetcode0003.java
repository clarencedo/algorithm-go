import java.util.HashMap;

class leetcode0003{
    public int lengthOfLongestSubstring(String s) {
        Integer ans = 0;
        HashMap<Character, Integer> window = new HashMap<>();
        Integer left = 0, right = 0;
        while (right < s.length()) {
            char rightChar = s.charAt(right);
            window.put(rightChar, window.getOrDefault(rightChar, 0) + 1);
            right++;
            while (window.get(rightChar) > 1) {
                char leftChar = s.charAt(left);
                window.put(leftChar, window.get(leftChar) - 1);
                left++;
            }

            ans = Math.max(ans, right - left);

        }
        return ans;
    }
}
