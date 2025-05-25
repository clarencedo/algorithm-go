import java.util.LinkedHashMap;
import java.util.LinkedList;

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

class LRUCacheII {
    private final int capacity;
    private final LinkedList<Integer> list;

    public LRUCacheII(int capacity) {
        this.capacity = capacity;
        this.list = new LinkedList<Integer>();
    }

    public void put(int key, int value) {
        if (list.size() < capacity) {
            list.addFirst(key);
        } else {
            list.removeLast();
            list.addFirst(key);
        }
    }

    public int get(int key) {
        var v = list.indexOf(key);
        if (v == -1) {
            return -1;
        }
        list.remove(v);
        list.addFirst(key);
        return v;
    }
}

class LRUListCache<E> {
    private final int maxSize;
    private final LinkedList<E> list = new LinkedList<>();

    public LRUListCache(int maxSize) {
        this.maxSize = maxSize;
    }

    public void add(E e) {
        if (list.size() < maxSize) {
            list.addFirst(e);
        } else {
            list.removeLast();
            list.addFirst(e);
        }
    }

    public E get(int index) {
        E e = list.get(index);
        list.remove(e);
        add(e);
        return e;
    }

    @Override
    public String toString() {
        return list.toString();
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */