# [57. 插入区间](https://leetcode.cn/problems/insert-interval/)

给你一个** 无重叠的** * ，* 按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

**示例 1：**

<pre><strong>输入：</strong>intervals = [[1,3],[6,9]], newInterval = [2,5]
<strong>输出：</strong>[[1,5],[6,9]]
</pre>

**示例 2：**

<pre><strong>输入：</strong>intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
<strong>输出：</strong>[[1,2],[3,10],[12,16]]
<strong>解释：</strong>这是因为新的区间 <code>[4,8]</code> 与 <code>[3,5],[6,7],[8,10]</code> 重叠。</pre>

**示例 3：**

<pre><strong>输入：</strong>intervals = [], newInterval = [5,7]
<strong>输出：</strong>[[5,7]]
</pre>

**示例 4：**

<pre><strong>输入：</strong>intervals = [[1,5]], newInterval = [2,3]
<strong>输出：</strong>[[1,5]]
</pre>

**示例 5：**

<pre><strong>输入：</strong>intervals = [[1,5]], newInterval = [2,7]
<strong>输出：</strong>[[1,7]]
</pre>

**提示：**

* `0 <= intervals.length <= 10<sup>4</sup>`
* `intervals[i].length == 2`
* `0 <= intervals[i][0] <= intervals[i][1] <= 10<sup>5</sup>`
* `intervals` 根据 `intervals[i][0]` 按 **升序** 排列
* `newInterval.length == 2`
* `0 <= newInterval[0] <= newInterval[1] <= 10<sup>5</sup>`
