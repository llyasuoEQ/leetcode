# [58. 最后一个单词的长度](https://leetcode.cn/problems/length-of-last-word/)

给你一个字符串 `s`，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 **最后一个** 单词的长度。

**单词** 是指仅由字母组成、不包含任何空格字符的最大子字符串。

**示例 1：**

<pre><strong>输入：</strong>s = "Hello World"
<strong>输出：</strong>5
<strong>解释：</strong>最后一个单词是“World”，长度为5。
</pre>

**示例 2：**

<pre><strong>输入：</strong>s = "   fly me   to   the moon  "
<strong>输出：</strong>4<strong>
解释：</strong>最后一个单词是“moon”，长度为4。
</pre>

**示例 3：**

<pre><strong>输入：</strong>s = "luffy is still joyboy"
<strong>输出：</strong>6
<strong>解释：</strong>最后一个单词是长度为6的“joyboy”。
</pre>

**提示：**

* `1 <= s.length <= 10<sup>4</sup>`
* `s` 仅有英文字母和空格 `' '` 组成
* `s` 中至少存在一个单词
