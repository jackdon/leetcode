---
id: bao-han-minhan-shu-de-zhan-lcof
title: 包含min函数的栈
---
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

 

**示例:**


<pre>MinStack minStack = new MinStack();<br/>minStack.push(-2);<br/>minStack.push(0);<br/>minStack.push(-3);<br/>minStack.min();   --&gt; 返回 -3.<br/>minStack.pop();<br/>minStack.top();      --&gt; 返回 0.<br/>minStack.min();   --&gt; 返回 -2.<br/></pre>

 

**提示：**

- 各函数的调用总次数不超过 20000 次
 

注意：本题与主站 155 题相同：[https://leetcode-cn.com/problems/min-stack/](https://leetcode-cn.com/problems/min-stack/)
