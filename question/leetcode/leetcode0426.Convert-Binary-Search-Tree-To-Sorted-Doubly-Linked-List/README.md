# [426.将二叉搜索树转化为排序的双向链表](https://leetcode.cn/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/description/)

# [426. Convert Binary Search Tree to Sorted Doubly Linked List 🔒](https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list)

## 题目

将一个 **二叉搜索树** 就地转化为一个 **已排序的双向循环链表** 。

对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

特别地，我们希望可以 **原地**完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。

**示例 1：**

    输入：root = [4,2,5,1,3]

![case1](../../../resources/leetcode/leetcode0426/1.png)

    输出：[1,2,3,4,5]

![case2](./../../../resources/leetcode/leetcode0426/2.png)

    解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。

**示例 2：**

    输入：root = [2,1,3]
    输出：[1,2,3]

**提示：**

- 树中节点的数量在范围 **[0, 2000]** 中
- `-1000 <= Node.val <= 1000`
  树中的所有值都是**独一无二**的