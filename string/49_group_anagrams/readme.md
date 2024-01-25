# [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/)

给你一个字符串数组，请你将 **字母异位词** 组合在一起。可以按任意顺序返回结果列表。

**字母异位词** 是由重新排列源单词的所有字母得到的一个新单词。

**示例 1:**

<pre><strong>输入:</strong> strs = <code>["eat", "tea", "tan", "ate", "nat", "bat"]</code>
<strong>输出: </strong>[["bat"],["nat","tan"],["ate","eat","tea"]]</pre>

**示例 2:**

<pre><strong>输入:</strong> strs = <code>[""]</code>
<strong>输出: </strong>[[""]]
</pre>

**示例 3:**

<pre><strong>输入:</strong> strs = <code>["a"]</code>
<strong>输出: </strong>[["a"]]</pre>

**提示：**

* `1 <= strs.length <= 10<sup>4</sup>`
* `0 <= strs[i].length <= 100`
* `strs[i]` 仅包含小写字母
