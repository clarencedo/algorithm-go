import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;

class Solution {
	// Map
	public boolean isAnagram(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}
		Map<Character, Integer> sMap = new HashMap<>();
		Map<Character, Integer> tMap = new HashMap<>();
		for (char ch : s.toCharArray()) {
			sMap.put(ch, sMap.getOrDefault(ch, 0) + 1);
		}
		for (char ch : t.toCharArray()) {
			tMap.put(ch, tMap.getOrDefault(ch, 0) + 1);
		}

		return sMap.equals(tMap);
	}

	// 排序
	public boolean isAnagramII(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}
		char[] sArr = s.toCharArray();
		char[] tArr = t.toCharArray();
		Arrays.sort(sArr);
		Arrays.sort(tArr);

		return Arrays.equals(sArr, tArr);
	}

	// 计数统计
	public boolean isAnagramIII(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}

		int[] count = new int[26];
		for (char ch : s.toCharArray()) {
			count[ch - 'a']++;
		}
		for (char ch : t.toCharArray()) {
			count[ch - 'a']--;
			if (count[ch - 'a'] < 0) {
				return false;
			}
		}

		return true;
	}

	// stream
	public boolean isAnagramIV(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}

		Map<Character, Long> sMap = s.chars().mapToObj(c -> (char) c)
				.collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
		Map<Character, Long> tMap = s.chars().mapToObj(c -> (char) c)
				.collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));

		return sMap.equals(tMap);

	}

	// s= 'aa' t='bb' 异或就不能用,这个只能算一个思路
	public boolean isAnagramV(String s, String t) {
		if (s.length() != t.length()) {
			return false;
		}
		int xor = 0;
		for (char ch : s.toCharArray()) {
			xor ^= ch;
		}
		for (char ch : t.toCharArray()) {
			xor ^= ch;
		}
		return xor == 0;
	}
}