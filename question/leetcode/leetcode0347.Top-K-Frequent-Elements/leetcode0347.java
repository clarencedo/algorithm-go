import java.security.KeyStore.Entry;
import java.util.*;

class Solution {
	public int[] topKFrequent(int[] nums, int k) {
		Map<Integer, Integer> freqMap = new HashMap<>();
		for (int num : nums) {
			freqMap.put(num, freqMap.getOrDefault(num, 0) + 1);
		}

		List<Map.Entry<Integer, Integer>> entries = new ArrayList<>(freqMap.entrySet());
		entries.sort((a, b) -> b.getKey() - a.getKey());

		// List<Integer> result = new ArrayList<>();
		int[] result = new int[k];
		for (int i = 0; i < k; i++) {
			result[i] = entries.get(i).getKey();
		}

		return result;
	}

	public int[] topKFrequentII(int[] nums, int k) {
		Map<Integer, Integer> freqMap = new HashMap<>();
		for (int num : nums) {
			freqMap.put(num, freqMap.getOrDefault(num, 0) + 1);
		}

		List<Integer>[] buckets = new ArrayList[nums.length + 1];
		for (int i = 0; i <= nums.length; i++) {
			buckets[i] = new ArrayList<>();
		}

		for (Map.Entry<Integer, Integer> entry : freqMap.entrySet()) {
			int num = entry.getKey();
			int freq = entry.getValue();
			buckets[freq].add(num);
		}

		List<Integer> result = new ArrayList<>();
		for (int i = buckets.length - 1; i >= 0 && result.size() < k; i--) {
			if (!buckets[i].isEmpty()) {
				result.addAll(buckets[i]);
			}
		}

		return result.stream().mapToInt(Integer::intValue).toArray();
	}

	public int[] topKFrequentIII(int[] nums, int k) {
		Map<Integer, Integer> freqMap = new HashMap<>();
		for (int num : nums) {
			freqMap.put(num, freqMap.getOrDefault(num, 0) + 1);
		}

		PriorityQueue<Map.Entry<Integer, Integer>> minHeap = new PriorityQueue<>(
				(a, b) -> a.getValue() - b.getValue());

		for (Map.Entry<Integer, Integer> entry : freqMap.entrySet()) {
			minHeap.offer(entry);
			if (minHeap.size() > k) {
				minHeap.poll();
			}
		}

		List<Integer> result = new ArrayList<>();
		while (!minHeap.isEmpty()) {
			result.add(minHeap.poll().getKey());
		}

		Collections.reverse(result);
		return result.stream().mapToInt(Integer::intValue).toArray();
	}
}