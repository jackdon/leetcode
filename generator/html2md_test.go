package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/TruthHun/html2md"
)

const htmlStr = "<p>Given a string <code>s</code>, return&nbsp;<em>the longest palindromic substring</em> in <code>s</code>.</p>\n\n<p>&nbsp;</p>\n<p><strong>Example 1:</strong></p>\n\n<pre>\n<strong>Input:</strong> s = &quot;babad&quot;\n<strong>Output:</strong> &quot;bab&quot;\n<strong>Note:</strong> &quot;aba&quot; is also a valid answer.\n</pre>\n\n<p><strong>Example 2:</strong></p>\n\n<pre>\n<strong>Input:</strong> s = &quot;cbbd&quot;\n<strong>Output:</strong> &quot;bb&quot;\n</pre>\n\n<p><strong>Example 3:</strong></p>\n\n<pre>\n<strong>Input:</strong> s = &quot;a&quot;\n<strong>Output:</strong> &quot;a&quot;\n</pre>\n\n<p><strong>Example 4:</strong></p>\n\n<pre>\n<strong>Input:</strong> s = &quot;ac&quot;\n<strong>Output:</strong> &quot;a&quot;\n</pre>\n\n<p>&nbsp;</p>\n<p><strong>Constraints:</strong></p>\n\n<ul>\n\t<li><code>1 &lt;= s.length &lt;= 1000</code></li>\n\t<li><code>s</code> consist of only digits and English letters (lower-case and/or upper-case),</li>\n</ul>\n"

func TestHtml2Md(t *testing.T) {
	mdStr := html2md.Convert(strings.ReplaceAll(htmlStr, "\n", "<br>"))
	ioutil.WriteFile("longest-palindromic-substring.md", []byte(mdStr), 0777)
	t.Log(mdStr)
}
