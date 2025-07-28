# OrderedMap

一个简单的 Go 有序映射实现，保留键的插入顺序并支持 HTML 转义。


## 特性
- 保留键值对的插入顺序
- 基本操作：Set、Get、Delete、Clear
- 字符串值的 HTML 转义（可配置）
- 自定义字符串格式化


## 安装
```bash
go get github.com/Patrick7241/orderedMap  
```


## 使用示例

```go
package main

import (
    "fmt"
    "github.com/Patrick7241/orderedMap"
)

func main() {
    // 创建一个新的 OrderedMap（默认启用 HTML 转义）
    om := orderedMap.New()

    // 添加键值对
    om.Set("name", "Alice")
    om.Set("age", 30)
    om.Set("bio", "<script>alert('hi')</script>")  // 默认会被转义

    // 获取值
    if val, exists := om.Get("name"); exists {
        fmt.Println("Name:", val)  // 输出: Name: Alice
    }

    // 打印映射（按插入顺序）
    fmt.Println(om)  // 输出: {"name": "Alice", "age": "30", "bio": "&lt;script&gt;alert(&#39;hi&#39;)&lt;/script&gt;"}

    // 禁用 HTML 转义
    om.SetEscapeHTML(false)
    fmt.Println(om)  // 输出: {"name": "Alice", "age": "30", "bio": "<script>alert('hi')</script>"}

    // 删除键
    om.Delete("age")
    fmt.Println(om)  // 输出: {"name": "Alice", "bio": "<script>alert('hi')</script>"}

    // 清空所有条目
    om.Clear()
    fmt.Println(om)  // 输出: {}
}
```


## API 说明

| 方法                | 描述                                      |
|---------------------|-------------------------------------------|
| `New(escapeHTML...)` | 创建新的 OrderedMap（默认启用 HTML 转义） |
| `Set(key, value)`   | 添加或更新键值对                          |
| `Get(key)`          | 通过键获取值（返回值和存在性标志）        |
| `Delete(key)`       | 删除指定键值对                            |
| `Clear()`           | 清空所有键值对                            |
| `SetEscapeHTML(on)` | 启用/禁用字符串值的 HTML 转义             |
| `String()`          | 获取字符串表示（按插入顺序）              |
