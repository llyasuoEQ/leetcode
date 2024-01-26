# [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)


给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

**示例 1：**

<pre><strong>输入：</strong>lists = [[1,4,5],[1,3,4],[2,6]]
<strong>输出：</strong>[1,1,2,3,4,4,5,6]
<strong>解释：</strong>链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
</pre>

**示例 2：**

<pre><strong>输入：</strong>lists = []
<strong>输出：</strong>[]
</pre>

**示例 3：**

<pre><strong>输入：</strong>lists = [[]]
<strong>输出：</strong>[]
</pre>

**提示：**

* `k == lists.length`
* `0 <= k <= 10^4`
* `0 <= lists[i].length <= 500`
* `-10^4 <= lists[i][j] <= 10^4`
* `lists[i]` 按 **升序** 排列
* `lists[i].length` 的总和不超过 `10^4`
