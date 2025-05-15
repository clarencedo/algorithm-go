from collections import OrderedDict


class LRUCache:

    def __init__(self, capacity: int):
        self.cache = OrderedDict()  # 存key-value
        self.capacity = capacity  # 缓存最大容量

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1

        self.cache.move_to_end(key, last=False)
        return self.cache[key]

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            self.cache.move_to_end(key, last=False)
            self.cache[key] = value
        else:
            if len(self.cache) == self.capacity:
                self.cache.popitem(last=True)
            self.cache[key] = value
            self.cache.move_to_end(key, last=False)
