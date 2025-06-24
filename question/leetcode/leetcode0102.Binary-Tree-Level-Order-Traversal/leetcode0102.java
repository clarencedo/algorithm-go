import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

class TreeNode {
	int val;
	TreeNode left;
	TreeNode right;

	TreeNode() {
	}

	TreeNode(int val) {
		this.val = val;
	}

	TreeNode(int val, TreeNode left, TreeNode right) {
		this.val = val;
		this.left = left;
		this.right = right;
	}
}

class Solution {
	public List<List<Integer>> levelOrder(TreeNode root) {
		if (root == null) {
			return new ArrayList<>();
		}
		List<List<Integer>> ans = new ArrayList<>();
		Queue<TreeNode> queue = new LinkedList<>();
		queue.offer(root);

		while (!queue.isEmpty()) {
			int size = queue.size();
			List<Integer> levelValues = new ArrayList<>();

			for (int i = 0; i < size; i++) {
				TreeNode currNode = queue.poll();
				levelValues.add(currNode.val);

				if (currNode.left != null) {
					queue.offer(currNode.left);
				}
				if (currNode.right != null) {
					queue.offer(currNode.right);
				}
			}

			ans.add(levelValues);
		}
		return ans;
	}
}