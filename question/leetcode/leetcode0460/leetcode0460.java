package question.leetcode.leetcode0460;

import java.util.HashMap;
import java.util.LinkedHashSet;

class LFUCache {

    private final int capacity;
    private final HashMap<Integer, Integer> keyToVal;// key to val
    private final HashMap<Integer, Integer> keyToFreq;// key to frequency
    private final HashMap<Integer, LinkedHashSet<Integer>> freqToKeys; // frequency -> keys (LRU order) 某个freq
                                                                       // 对应的该频率下所有key的集合
    private int minFreq;// 当前最小频率

    public LFUCache(int capacity) {
        this.capacity = capacity;
        this.keyToVal = new HashMap<>();
        this.keyToFreq = new HashMap<>();
        this.freqToKeys = new HashMap<>();
        this.minFreq = 0;
    }

    public int get(int key) {
        if (!keyToVal.containsKey(key)) {
            return -1;
        }

        increaseFreq(key);
        return keyToVal.get(key);
    }

    public void put(int key, int value) {
        if (capacity <= 0) {
            return;
        }

        if (keyToVal.containsKey(key)) {
            keyToVal.put(key, value);
            increaseFreq(key);
            return;
        }

        if (keyToVal.size() >= capacity) {
            removeMinFreqKey();
        }

        keyToVal.put(key, value);
        keyToFreq.put(key, 1);
        freqToKeys.computeIfAbsent(1, k -> new LinkedHashSet<>()).add(key);
        // 新插入的key 频率为1
        minFreq = 1;
    }

    private void increaseFreq(int key) {
        int freq = keyToFreq.get(key);
        keyToFreq.put(key, freq + 1);

        // 从旧的频率的集合中移除
        freqToKeys.get(freq).remove(key);
        if (freqToKeys.get(freq).isEmpty()) {
            freqToKeys.remove(freq);
            if (freq == minFreq) {
                minFreq = freq + 1;
            }
        }

        // 加入新的频率的集合
        freqToKeys.computeIfAbsent(freq + 1, k -> new LinkedHashSet<>()).add(key);
    }

    // 淘汰minFreq的Key
    private void removeMinFreqKey() {
        LinkedHashSet<Integer> keyList = freqToKeys.get(minFreq);
        int deletedKey = keyList.iterator().next();// 取最早插入的key(LRU)
        keyList.remove(deletedKey);
        if (keyList.isEmpty()) {
            freqToKeys.remove(minFreq);
        }
        keyToVal.remove(deletedKey);
        keyToFreq.remove(deletedKey);
    }
}
