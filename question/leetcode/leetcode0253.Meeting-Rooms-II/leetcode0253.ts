function minMeetingRooms(intervals: number[][]): number {
	if (intervals.length === 0) {
		return 0;
	}

	//按开始时间排序
	intervals.sort((a, b) => a[0] - b[0]);

	//模拟最小堆,会议结束时间
	const heap: number[] = []

	for (const [start, end] of intervals) {
		//如果当前会议开始时间 >= 最小结束时间，可以复用会议室
		if (heap.length > 0 && start >= heap[0]) {
			//移除最早结束的会议
			heap.shift();
		}
		//加入当前会议的结束时间
		heap.push(end);

		//维持headp为最小堆结构
		heap.sort((a, b) => a - b);

	}

	return heap.length;
};