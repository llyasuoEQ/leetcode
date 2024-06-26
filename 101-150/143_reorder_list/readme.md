# [143. 重排链表](https://leetcode.cn/problems/reorder-list/)

给定一个单链表 `L` 的头节点 `head` ，单链表 `L` 表示为：

```
L0 → L1 → … → Ln - 1 → Ln
```

请将其重新排列后变为：

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```

不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

**示例 1：**

![](https://pic.leetcode-cn.com/1626420311-PkUiGI-image.png)

<pre><strong>输入：</strong>head = [1,2,3,4]
<strong>输出：</strong>[1,4,2,3]</pre>

**示例 2：**

![](https://pic.leetcode-cn.com/1626420320-YUiulT-image.png)

<pre><strong>输入：</strong>head = [1,2,3,4,5]
<strong>输出：</strong>[1,5,2,4,3]</pre>

**提示：**

* 链表的长度范围为 `[1, 5 * 10<sup>4</sup>]`
* `1 <= node.val <= 1000`
