# [61. 旋转链表](https://leetcode.cn/problems/rotate-list/)

给你一个链表的头节点 `head` ，旋转链表，将链表每个节点向右移动 `k` 个位置。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/11/13/rotate1.jpg)

<pre><strong>输入：</strong>head = [1,2,3,4,5], k = 2
<strong>输出：</strong>[4,5,1,2,3]
</pre>

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/11/13/roate2.jpg)

<pre><strong>输入：</strong>head = [0,1,2], k = 4
<strong>输出：</strong>[2,0,1]
</pre>

**提示：**

* 链表中节点的数目在范围 `[0, 500]` 内
* `-100 <= Node.val <= 100`
* `0 <= k <= 2 * 10<sup>9</sup>`
