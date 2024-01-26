# [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)

给你一个整数数组 `nums` ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**子数组 **是数组中的一个连续部分。

**示例 1：**

<pre><strong>输入：</strong>nums = [-2,1,-3,4,-1,2,1,-5,4]
<strong>输出：</strong>6
<strong>解释：</strong>连续子数组 [4,-1,2,1] 的和最大，为 6 。
</pre>

**示例 2：**

<pre><strong>输入：</strong>nums = [1]
<strong>输出：</strong>1
</pre>

**示例 3：**

<pre><strong>输入：</strong>nums = [5,4,-1,7,8]
<strong>输出：</strong>23
</pre>

**提示：**

* `1 <= nums.length <= 10<sup>5</sup>`
* `-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup>`

 **进阶：** 如果你已经实现复杂度为 `O(n)` 的解法，尝试使用更为精妙的 **分治法** 求解。
