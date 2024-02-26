# [62. 不同路径](https://leetcode.cn/problems/unique-paths/)

一个机器人位于一个 `m x n` 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

**示例 1：**

![](https://pic.leetcode.cn/1697422740-adxmsI-image.png)

<pre><strong>输入：</strong>m = 3, n = 7
<strong>输出：</strong>28</pre>

**示例 2：**

<pre><strong>输入：</strong>m = 3, n = 2
<strong>输出：</strong>3
<strong>解释：</strong>
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
</pre>

**示例 3：**

<pre><strong>输入：</strong>m = 7, n = 3
<strong>输出：</strong>28
</pre>

**示例 4：**

<pre><strong>输入：</strong>m = 3, n = 3
<strong>输出：</strong>6</pre>

**提示：**

* `1 <= m, n <= 100`
* 题目数据保证答案小于等于 `2 * 10<sup>9</sup>`
