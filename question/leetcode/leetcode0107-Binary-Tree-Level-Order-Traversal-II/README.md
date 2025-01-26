# [107.二叉树的层序遍历II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
# [107.Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)

## 题目

给你二叉树的根节点 `root` ，返回其节点值 **自底向上的层序遍历** 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）



**示例 1：**

   [![case1](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

    输入：root = [3,9,20,null,null,15,7]
    输出：[[15,7],[9,20],[3]]

**示例 2：**

    输入：root = [1]
    输出：[[1]]

**示例 3：**

    输入：root = []
    输出：[]


**提示：**

- 树中节点数目在范围 `[0, 2000]` 内
- `-1000 <= Node.val <= 1000`


## 题目大意

按层序从下到上遍历一颗树。

## 解题思路

用一个队列即可实现。



