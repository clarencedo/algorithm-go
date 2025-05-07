class TreeNode {
	constructor(val = 0, left = null, right = null) {
		this.val = val;
		this.left = left;
		this.right = right;
	}
}

class CBTInserter {
	constructor(root) {
		this.root = root;
		this.q = [];

		const queue = [];
		if (root !== null) {
			queue.push(root);
		}
		while (queue.length > 0) {
			const node = queue.shift();
			if (node.left !== null) {
				queue.push(node.left);
			}
			if (node.right !== null) {
				queue.push(node.right);
			}
			if (node.left === null || node.right === null) {
				this.q.push(node);
			}
		}
	}

	insert(val) {
		const parent = this.q[0];
		const newNode = new TreeNode(val);
		if (parent.left === null) {
			parent.left = newNode;
		} else {
			parent.right = newNode;
			this.q.shift();
		}
		this.q.push(newNode);
		return parent.val;
	}

	getRoot() {
		return this.root;
	}
}
