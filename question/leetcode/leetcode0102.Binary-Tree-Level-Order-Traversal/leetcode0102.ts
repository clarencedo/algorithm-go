class TreeNode {
	val: number
	left: TreeNode | null
	right: TreeNode | null
	constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.left = (left === undefined ? null : left)
		this.right = (right === undefined ? null : right)
	}
}

function levelOrder(root: TreeNode | null): number[][] {
	if (!root) { return [] };

	const ans: number[][] = [];
	let queue: TreeNode[] = [root];

	while (queue.length > 0) {
		const levelValues: number[] = []
		const nextQueue: TreeNode[] = [];

		for (const node of queue) {
			levelValues.push(node.val);
			if (node.left) nextQueue.push(node.left);
			if (node.right) nextQueue.push(node.right);
		}

		ans.push(levelValues);
		queue = nextQueue;
	}

	return ans;
};

function levelOrderII(root: TreeNode | null): number[][] {
	if (!root) { return [] };

	const ans: number[][] = [];
	let queue: TreeNode[] = [root];

	while (queue.length > 0) {
		const size = queue.length;
		const levelValues: number[] = [];
		for (let i = 0; i < size; i++) {
			const currentNode = queue.shift();
			levelValues.push(currentNode!.val)
			if (currentNode!.left) {
				queue.push(currentNode!.left);
			}
			if (currentNode!.right) {
				queue.push(currentNode!.right);
			}
		}
		ans.push(levelValues);
	}

	return ans;
};
