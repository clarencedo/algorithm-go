import java.util.List;
import java.util.Set;
import java.util.HashSet;

class Solution {
	public boolean wordBreak(String s, List<String> wordDict) {
		Set<String> wordSet = new HashSet<>(wordDict);
		boolean[] dp = new boolean[s.length() + 1];
		dp[0] = true;

		for (int i = 0; i <= s.length(); i++) {
			for (int j = 0; j < i; j++) {
				String word = s.substring(j, i);
				if (dp[j] && wordSet.contains(word)) {
					dp[i] = true;
					break;
				}

			}
		}
		return dp[s.length()];
	}
}