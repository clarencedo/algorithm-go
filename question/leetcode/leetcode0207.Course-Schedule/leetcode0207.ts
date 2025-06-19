//基于 BFS 拓扑排序 - Kahn’s Algorithm
//本质上是判断一个有向图是否有环,有环就无法满足
function canFinish(numCourses: number, prerequisites: number[][]): boolean {
	const inDegree: number[] = new Array(numCourses).fill(0); // 入度表
	const graph: Map<number, number[]> = new Map(); // 邻接表

	// 构建图和入度表
	for (const [to, from] of prerequisites) {
		inDegree[to]++; // to课程入度+1

		if (!graph.has(from)) {
			graph.set(from, []);
		}
		graph.get(from)!.push(to); // from -> to
	}

	// 将入度为0的课程加入队列
	const queue: number[] = [];
	for (let i = 0; i < numCourses; i++) {
		if (inDegree[i] === 0) {
			queue.push(i);
		}
	}

	let finished = 0;

	while (queue.length > 0) {
		const course = queue.shift()!;
		finished++;

		for (const next of graph.get(course) || []) {
			inDegree[next]--;
			if (inDegree[next] === 0) {
				queue.push(next);
			}
		}
	}

	return finished === numCourses;
}
