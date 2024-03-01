# [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)

给定一个含有 `n` 个正整数的数组和一个正整数 `target`** 。**

找出该数组中满足其总和大于等于 `target` 的长度最小的 **连续子数组** `[nums<sub>l</sub>, nums<sub>l+1</sub>, ..., nums<sub>r-1</sub>, nums<sub>r</sub>]` ，并返回其长度 **。** 如果不存在符合条件的子数组，返回 `0` 。

**示例 1：**

<pre><strong>输入：</strong>target = 7, nums = [2,3,1,2,4,3]
<strong>输出：</strong>2
<strong>解释：</strong>子数组 <code>[4,3]</code> 是该条件下的长度最小的子数组。
</pre>

**示例 2：**

<pre><strong>输入：</strong>target = 4, nums = [1,4,4]
<strong>输出：</strong>1
</pre>

**示例 3：**

<pre><strong>输入：</strong>target = 11, nums = [1,1,1,1,1,1,1,1]
<strong>输出：</strong>0
</pre>

**提示：**

* `1 <= target <= 10<sup>9</sup>`
* `1 <= nums.length <= 10<sup>5</sup>`
