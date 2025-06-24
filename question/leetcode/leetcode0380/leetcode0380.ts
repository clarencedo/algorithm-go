class RandomizedSet {
	private map: Map<number, number>; // val -> index
	private arr: number[]; // 存储值

	constructor() {
		this.map = new Map();
		this.arr = [];
	}

	insert(val: number): boolean {
		if (!this.map.has(val)) {
			this.arr.push(val);
			this.map.set(val, this.arr.length - 1);
			return true;
		}

		return false;
	}

	//数组中删除一个元素的时间复杂度不是O(1)
	//所以需要 先 换大欧子后 再pop
	remove(val: number): boolean {
		if (!this.map.has(val)) { return false; }

		const idx = this.map.get(val)!;
		const lastVal = this.arr[this.arr.length - 1]

		//覆盖idx为止的元素
		this.arr[idx] = lastVal;
		this.map.set(lastVal, idx);

		//删除最后一个元素
		this.arr.pop();
		this.map.delete(val);

		return true;
	}

	getRandom(): number {
		const idx = Math.floor(Math.random() * this.arr.length);
		return this.arr[idx];
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * var obj = new RandomizedSet()
 * var param_1 = obj.insert(val)
 * var param_2 = obj.remove(val)
 * var param_3 = obj.getRandom()
 */