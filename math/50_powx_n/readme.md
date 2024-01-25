# [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)

实现 [pow( *x* ,  *n* )](https://www.cplusplus.com/reference/valarray/pow/) ，即计算 `x` 的整数 `n` 次幂函数（即，`x<sup>n</sup>` ^ ^ ）。

**示例 1：**

<pre><strong>输入：</strong>x = 2.00000, n = 10
<strong>输出：</strong>1024.00000
</pre>

**示例 2：**

<pre><strong>输入：</strong>x = 2.10000, n = 3
<strong>输出：</strong>9.26100
</pre>

**示例 3：**

<pre><strong>输入：</strong>x = 2.00000, n = -2
<strong>输出：</strong>0.25000
<strong>解释：</strong>2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25
</pre>

**提示：**

* `-100.0 < x < 100.0`
* `-2<sup>31</sup> <= n <= 2<sup>31</sup>-1`
* `n` 是一个整数
* 要么 `x` 不为零，要么 `n > 0` 。
