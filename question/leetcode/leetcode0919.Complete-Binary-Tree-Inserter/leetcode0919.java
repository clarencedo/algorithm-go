import java.util.ArrayDeque;
import java.util.ArrayList;
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

class CBTInserter {
    TreeNode root;
    List<TreeNode> q;

    public CBTInserter(TreeNode root) {
        this.root = root;
        this.q = new ArrayList<>();

        Queue<TreeNode> queue = new ArrayDeque<>();
        if (root != null) {
            queue.offer(root);
        }
        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            if (node.left != null) {
                queue.offer(node.left);
            }
            if (node.right != null) {
                queue.offer(node.right);
            }
            if (node.left == null || node.right == null) {
                q.add(node);
            }
        }
    }

    public int insert(int val) {
        TreeNode parent = q.get(0);
        TreeNode newNode = new TreeNode(val);
        if (parent.left == null) {
            parent.left = newNode;
        } else {
            parent.right = newNode;
            q.remove(0);
        }
        q.add(newNode);
        return parent.val;
    }

    public TreeNode get_root() {
        return this.root;

    }
}
