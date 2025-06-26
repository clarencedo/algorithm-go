import java.util.HashMap;
import java.util.Map;

class Solution {
	public boolean canConstruct(String ransomNote, String magazine) {
		Map<Character, Integer> magazineMp = new HashMap<>();

		for (char ch : magazine.toCharArray()) {
			magazineMp.put(ch, magazineMp.getOrDefault(ch, 0) + 1);
		}

		for (char ch : ransomNote.toCharArray()) {
			if (!magazineMp.containsKey(ch) || magazineMp.get(ch) <= 0) {
				return false;
			}
			magazineMp.put(ch, magazineMp.get(ch) - 1);
		}

		return true;
	}
}