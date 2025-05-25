# 实现方式1：使用OrderedDict
from collections import OrderedDict

class LRUCache_OrderedDict:
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


# 实现方式2：使用普通字典和双向链表
class DLinkedNode:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None

class LRUCache:
    def __init__(self, capacity: int):
        self.cache = {}  # 哈希表，存储key到节点的映射
        self.capacity = capacity  # 缓存最大容量
        self.size = 0  # 当前缓存大小
        
        # 使用伪头部和伪尾部节点
        self.head = DLinkedNode()  # 最近使用的放在头部
        self.tail = DLinkedNode()  # 最久未使用的放在尾部
        self.head.next = self.tail
        self.tail.prev = self.head
    
    # 将节点添加到双向链表的头部
    def add_to_head(self, node):
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node
    
    # 从双向链表中删除一个节点
    def remove_node(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
    
    # 将节点移到头部（表示最近使用）
    def move_to_head(self, node):
        self.remove_node(node)
        self.add_to_head(node)
    
    # 删除尾部节点（最久未使用的）
    def remove_tail(self):
        node = self.tail.prev
        self.remove_node(node)
        return node

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        
        # 如果key存在，先通过哈希表定位，再移到头部
        node = self.cache[key]
        self.move_to_head(node)
        return node.value

    def put(self, key: int, value: int) -> None:
        if key not in self.cache:
            # 如果key不存在，创建一个新的节点
            node = DLinkedNode(key, value)
            # 添加进哈希表
            self.cache[key] = node
            # 添加至双向链表的头部
            self.add_to_head(node)
            self.size += 1
            
            # 如果超出容量，删除双向链表的尾部节点
            if self.size > self.capacity:
                # 删除双向链表的尾部节点
                removed = self.remove_tail()
                # 删除哈希表中对应的项
                del self.cache[removed.key]
                self.size -= 1
        else:
            # 如果key存在，先通过哈希表定位，再修改value，并移到头部
            node = self.cache[key]
            node.value = value
            self.move_to_head(node)
