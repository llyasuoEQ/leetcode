# [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/)


将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

<pre><strong>输入：</strong>l1 = [1,2,4], l2 = [1,3,4]
<strong>输出：</strong>[1,1,2,3,4,4]
</pre>

**示例 2：**

<pre><strong>输入：</strong>l1 = [], l2 = []
<strong>输出：</strong>[]
</pre>

**示例 3：**

<pre><strong>输入：</strong>l1 = [], l2 = [0]
<strong>输出：</strong>[0]
</pre>

**提示：**

* 两个链表的节点数目范围是 `[0, 50]`
* `-100 <= Node.val <= 100`
* `l1` 和 `l2` 均按 **非递减顺序** 排列
