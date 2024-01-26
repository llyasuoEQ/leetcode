# [44. 通配符匹配](https://leetcode.cn/problems/wildcard-matching/)

给你一个输入字符串 (`s`) 和一个字符模式 (`p`) ，请你实现一个支持 `'?'` 和 `'*'` 匹配规则的通配符匹配：

* `'?'` 可以匹配任何单个字符。
* `'*'` 可以匹配任意字符序列（包括空字符序列）。

判定匹配成功的充要条件是：字符模式必须能够 **完全匹配** 输入字符串（而不是部分匹配）。

 **示例 1：**

<pre><strong>输入：</strong>s = "aa", p = "a"
<strong>输出：</strong>false
<strong>解释：</strong>"a" 无法匹配 "aa" 整个字符串。
</pre>

**示例 2：**

<pre><strong>输入：</strong>s = "aa", p = "*"
<strong>输出：</strong>true
<strong>解释：</strong>'*' 可以匹配任意字符串。
</pre>

**示例 3：**

<pre><strong>输入：</strong>s = "cb", p = "?a"
<strong>输出：</strong>false
<strong>解释：</strong>'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
</pre>

**提示：**

* `0 <= s.length, p.length <= 2000`
* `s` 仅由小写英文字母组成
* `p` 仅由小写英文字母、`'?'` 或 `'*'` 组成
