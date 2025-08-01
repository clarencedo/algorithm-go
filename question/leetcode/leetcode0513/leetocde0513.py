# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findBottomLeftValue(self, root: Optional[TreeNode]) -> int:
        ''' 
        bfs，从右往左入队列， 这样队列清空之后最后一个节点就是最左边的节点
        '''
        node = root
        q = [node]
        while len(q) > 0:
            node, q = q[0], q[1:]
            if node.right:
                q.append(node.right)
            if node.left:
                q.append(node.left)

        return node.val
