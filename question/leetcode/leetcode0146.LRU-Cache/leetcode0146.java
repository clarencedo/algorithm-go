import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.Map;

class LRUCache {
    private final int capacity;
    private final LinkedHashMap<Integer, Integer> cache;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.cache = new LinkedHashMap<Integer, Integer>(capacity, 0.75f, true) {
            @Override
            protected boolean removeEldestEntry(Map.Entry<Integer, Integer> eldest) {
                return size() > capacity;
            }
        };
    }

    public int get(int key) {
        return cache.getOrDefault(key, -1);
    }

    public void put(int key, int value) {
        cache.put(key, value);
    }
}

class LRU {
    private final int capacity;
    private final LinkedList<Integer> list;
    private final HashMap<Integer, Integer> map;

    public LRU(int capacity) {
        this.capacity = capacity;
        this.list = new LinkedList<Integer>();
        this.map = new HashMap<Integer, Integer>();
    }

    public int get(int key) {
        if (!map.containsKey(key)) {
            return -1;
        }

        list.remove((Integer) key);
        list.addFirst(key);
        return map.get(key);
    }

    public void put(int key, int val) {
        if (map.containsKey(key)) {
            list.remove((Integer) key);
        } else if (list.size() == capacity) {
            int oldKey = list.removeLast();
            map.remove(oldKey);
        }

        list.addFirst(key);
        map.put(key, val);
    }

}
