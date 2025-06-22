class LinkedNode {
	key: number;
	val: number;
	prev: LinkedNode | null = null;
	next: LinkedNode | null = null;

	constructor(key?: number, val?: number) {
		this.key = key ?? 0;
		this.val = val ?? 0;
	}
}
class LRUCache {
	private capacity: number;
	private cache: Map<number, LinkedNode>;
	private head: LinkedNode;
	private tail: LinkedNode;

	constructor(capacity: number) {
		this.capacity = capacity;
		this.cache = new Map();

		this.head = new LinkedNode();
		this.tail = new LinkedNode();
		this.head.next = this.tail;
		this.tail.prev = this.head;
	}

	get(key: number): number {
		const node = this.cache.get(key);
		if (!node) {
			return -1
		}

		this.moveToHead(node);
		return node.val;
	}

	put(key: number, value: number): void {
		const node = this.cache.get(key);
		if (node) {
			node.val = value;
			this.moveToHead(node);
		} else {
			const newNode = new LinkedNode(key, value);
			this.cache.set(key, newNode);
			this.addToHead(newNode);

			if (this.cache.size > this.capacity) {
				const tailNode = this.removeTail();
				this.cache.delete(tailNode.key);
			}
		}

	}
	private addToHead(node: LinkedNode) {
		node.prev = this.head;
		node.next = this.head.next;
		this.head.next!.prev = node;
		this.head.next = node;
	}

	private removeNode(node: LinkedNode) {
		node.prev!.next = node.next;
		node.next!.prev = node.prev;
	}

	private moveToHead(node: LinkedNode) {
		this.removeNode(node);
		this.addToHead(node);
	}

	private removeTail(): LinkedNode {
		const node = this.tail.prev!;
		this.removeNode(node);
		return node;
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */