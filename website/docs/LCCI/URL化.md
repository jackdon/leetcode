---
id: string-to-url-lcci
title: URL化
---
URL化。编写一种方法，将字符串中的空格全部替换为<code>%20</code>。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用<code>Java</code>实现的话，请使用字符数组实现，以便直接在数组上操作。）

 

**示例 1：**


<pre><br/><strong>输入</strong>：&#34;Mr John Smith    &#34;, 13<br/><strong>输出</strong>：&#34;Mr%20John%20Smith&#34;<br/></pre>

**示例 2：**


<pre><br/><strong>输入</strong>：&#34;               &#34;, 5<br/><strong>输出</strong>：&#34;%20%20%20%20%20&#34;<br/></pre>

 

**提示：**


- 字符串长度在 [0, 500000] 范围内。
