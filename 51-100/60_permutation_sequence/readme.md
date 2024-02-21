# [60. 排列序列](https://leetcode.cn/problems/permutation-sequence/)

给出集合 `[1,2,3,...,n]`，其所有元素共有 `n!` 种排列。

按大小顺序列出所有排列情况，并一一标记，当 `n = 3` 时, 所有排列如下：

1. `"123"`
2. `"132"`
3. `"213"`
4. `"231"`
5. `"312"`
6. `"321"`

给定 `n` 和 `k`，返回第 `k` 个排列。

**示例 1：**

<pre><strong>输入：</strong>n = 3, k = 3
<strong>输出：</strong>"213"
</pre>

**示例 2：**

<pre><strong>输入：</strong>n = 4, k = 9
<strong>输出：</strong>"2314"
</pre>

**示例 3：**

<pre><strong>输入：</strong>n = 3, k = 1
<strong>输出：</strong>"123"
</pre>

**提示：**

* `1 <= n <= 9`
* `1 <= k <= n!`
