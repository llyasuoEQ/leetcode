# [29. 两数相除](https://leetcode.cn/problems/divide-two-integers/)

给你两个整数，被除数 `dividend` 和除数 `divisor`。将两数相除，要求 **不使用** 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（`truncate`）其小数部分。例如，`8.345` 将被截断为 `8` ，`-2.7335` 将被截断至 `-2` 。

返回被除数 `dividend` 除以除数 `divisor` 得到的 **商** 。

 **注意：** 假设我们的环境只能存储 **32 位** 有符号整数，其数值范围是 `[−2<sup>31</sup>,  2<sup>31 </sup>− 1]` 。本题中，如果商 **严格大于** `2<sup>31 </sup>− 1` ，则返回 `2<sup>31 </sup>− 1` ；如果商 **严格小于** `-2<sup>31</sup>` ，则返回 `-2<sup>31</sup>` ^ ^ 。

**示例 1:**

<pre><strong>输入:</strong> dividend = 10, divisor = 3
<strong>输出:</strong> 3
<strong>解释: </strong>10/3 = 3.33333.. ，向零截断后得到 3 。</pre>

**示例 2:**

<pre><strong>输入:</strong> dividend = 7, divisor = -3
<strong>输出:</strong> -2
<strong>解释:</strong> 7/-3 = -2.33333.. ，向零截断后得到 -2 。</pre>

**提示：**

* `-2<sup>31</sup> <= dividend, divisor <= 2<sup>31</sup> - 1`
* `divisor != 0`
