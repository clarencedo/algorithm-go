# [445.两数相加](https://leetcode.cn/problems/add-two-numbers-ii/description/)

# [445. Add Two Numbers II](https://leetcode.com/problems/add-two-numbers-ii/)

## 题目

给你两个**非空**链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

**示例 1：**

![case1](../../../resources/leetcode/leetcode0445/1.png)

    输入：l1 = [7,2,4,3], l2 = [5,6,4]
    输出：[7,8,0,7]

**示例 2：**

    输入：l1 = [0], l2 = [0]
    输出：[0]

**示例 3：**

    输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    输出：[8,9,9,9,0,0,0,1]

**提示：**

- 树中节点的数量在范围 **[0, 2000]** 中
- `-1000 <= Node.val <= 1000`
  树中的所有值都是**独一无二**的

- 每个链表中的节点数在范围 `[1, 100]` 内
- `0 <= Node.val <= 9`
- 题目数据保证列表表示的数字不含前导零

## 解题思路

- 本题是[2.两数相加](https://leetcode-cn.com/problems/add-two-numbers/)的变种，需要先将两个链表反转，然后再进行相加操作
- [206.反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)