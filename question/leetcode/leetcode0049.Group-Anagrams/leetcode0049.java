import java.util.List;
import java.util.Map;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> mp = new HashMap<>();
        for (String s : strs) {
            char[] cs = s.toCharArray();
            Arrays.sort(cs);
            String sortedStr = new String(cs);
            if (mp.containsKey(sortedStr)) {
                mp.get(sortedStr).add(s);
            } else {
                List<String> list = new ArrayList<>();
                list.add(s);
                mp.put(sortedStr, list);
            }

        }
        List<List<String>> ans = new ArrayList<>();

        for (String k : mp.keySet()) {
            ans.add(mp.get(k));
        }

        return ans;
    }
}