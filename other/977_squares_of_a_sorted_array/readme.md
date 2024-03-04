# [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/)

给你一个按 **非递减顺序** 排序的整数数组 `nums`，返回 **每个数字的平方** 组成的新数组，要求也按 **非递减顺序** 排序。

**示例 1：**

<pre><strong>输入：</strong>nums = [-4,-1,0,3,10]
<strong>输出：</strong>[0,1,9,16,100]
<strong>解释：</strong>平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]</pre>

**示例 2：**

<pre><strong>输入：</strong>nums = [-7,-3,2,3,11]
<strong>输出：</strong>[4,9,9,49,121]
</pre>

**提示：**

* `<span>1 <= nums.length <= </span>10<sup>4</sup>`
* `-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup>`
* `nums` 已按 **非递减顺序** 排序

**进阶：**

* 请你设计时间复杂度为 `O(n)` 的算法解决本问题
