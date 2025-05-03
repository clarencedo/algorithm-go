class Solution {
    public String longestCommonPrefix(String[] strs) {
        int n = strs.length;
        if (n == 0) {
            return "";
        }
        String prefix = "";
        String firstStr = strs[0];
        for (int i = 0; i < firstStr.length(); i++) {
            char ch = firstStr.charAt(i);
            for (int j = 1; j < strs.length; j++) {
                if (i >= strs[j].length() || strs[j].charAt(i) != ch) {
                    return prefix;
                }
            }
            prefix += ch;
        }

        return prefix;
    }

}
