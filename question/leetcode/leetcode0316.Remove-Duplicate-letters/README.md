# [316.去除重复字母](https://leetcode-cn.com/problems/remove-duplicate-letters/)

## 题目

给定一个二叉树，找到其中最大的二叉搜索树（BST）`子树`，并返回该
子树
的大小。其中，最大指的是子树节点数最多的。

**二叉搜索树(BST)** 中的所有节点都具备以下属性：

左子树的值小于其父（根）节点的值。

右子树的值大于其父（根）节点的值。

**注意**：子树必须包含其所有后代

**示例 1：**

![case1](./../../../resources/leetcode/leetcode0333/1.jpg)

    输入：root = [1,2,3,4,5]
    输出：[4,5,2,null,null,3,1]

**示例 2：**

    输入：root = []
    输出：[]

**示例 3：**

    输入：root = [1]
    输出：[1]

**提示：**

- 树上节点数目的范围是 `[0, 104]`
- `-104 <= Node.val <= 104`

进阶: 你能想出 `O(n)` 时间复杂度的解法吗？

## 解题思路

- 如何去掉一个字符，使得剩下的字符串字典序最小?

  - 找出最小的满足`s[i] > s[i+1]`的下标i，并去掉`s[i]`

- 如何保证去掉的字符在原字符串中存在？
  - 在考虑字符`s[i]`时，如果它已经存在于栈中，则不能删除它，否则会破坏栈的单调性。
  - 在弹出栈顶字符时，如果字符串在后面的位置上再也没有这一字符，则不能弹出栈顶字符。为此，需要记录每个字符的剩余数量，当这个值为 0 时，就不能弹出栈顶字符了。

## 代码

```go

package leetcode

func removeDuplicateLetters(s string) string {
    left := [26]int{}
    for _, ch := range s {
        left[ch-'a']++
    }
    stack := []byte{}
    inStack := [26]bool{}
    for i := range s {
        ch := s[i]
        if !inStack[ch-'a'] {
            for len(stack) > 0 && ch < stack[len(stack)-1] {
                last := stack[len(stack)-1] - 'a'
                if left[last] == 0 {
                    break
                }
                stack = stack[:len(stack)-1]
                inStack[last] = false
            }
            stack = append(stack, ch)
            inStack[ch-'a'] = true
        }
        left[ch-'a']--
    }
    return string(stack)
}

```
