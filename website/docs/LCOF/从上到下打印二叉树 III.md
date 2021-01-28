---
id: cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
title: 从上到下打印二叉树 III
---
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

 

例如:给定二叉树: <code>[3,9,20,null,null,15,7]</code>,


<pre>    3<br/>   / \<br/>  9  20<br/>    /  \<br/>   15   7<br/></pre>

返回其层次遍历结果：


<pre>[<br/>  [3],<br/>  [20,9],<br/>  [15,7]<br/>]<br/></pre>

 

**提示：**

- <code>节点总数 &lt;= 1000</code>