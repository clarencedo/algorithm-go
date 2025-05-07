using System;
using System.Collections.Generic;

class TreeNode
{
	int val;
	TreeNode left;
	TreeNode right;

	TreeNode() { }

	TreeNode(int val)
	{
		this.val = val;
	}

	TreeNode(int val, TreeNode left, TreeNode right)
	{
		this.val = val;
		this.left = left;
		this.right = right;
	}
}

class CBTInserter
{
	private TreeNode root;
	private List<TreeNode> q;

	CBTInserter(TreeNode root)
	{
		this.root = root;
		this.q = new List<TreeNode>();

		Queue<TreeNode> queue = new Queue<TreeNode>();
		if (root != null)
		{
			queue.Enqueue(root);
		}
		while (queue.Count > 0)
		{
			TreeNode node = queue.Dequeue();
			if (node.left != null)
			{
				queue.Enqueue(node.left);
			}
			if (node.right != null)
			{
				queue.Enqueue(node.right);
			}
			if (node.left == null || node.right == null)
			{
				q.Add(node);
			}
		}
	}

	int Insert(int val)
	{
		TreeNode parent = q[0];
		TreeNode newNode = new TreeNode(val);
		if (parent.left == null)
		{
			parent.left = newNode;
		}
		else
		{
			parent.right = newNode;
			q.RemoveAt(0);
		}
		q.Add(newNode);
		return parent.val;
	}

	TreeNode GetRoot()
	{
		return this.root;
	}
}