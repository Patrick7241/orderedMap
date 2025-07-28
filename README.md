# OrderedMap

A simple ordered map implementation for Go, which preserves the insertion order of keys and supports HTML escaping.


## Features
- Preserves insertion order of key-value pairs
- Basic operations: Set, Get, Delete, Clear
- HTML escaping for string values (configurable)
- Custom string representation


## Installation
```bash
go get github.com/Patrick7241/orderedMap  
```


## Usage Example

```go
package main

import (
    "fmt"
    "github.com/Patrick7241/orderedMap"  
)

func main() {
    // Create a new OrderedMap (HTML escaping enabled by default)
    om := orderedmap.New()

    // Add key-value pairs
    om.Set("name", "Alice")
    om.Set("age", 30)
    om.Set("bio", "<script>alert('hi')</script>")  // Will be escaped by default

    // Get a value
    if val, exists := om.Get("name"); exists {
        fmt.Println("Name:", val)  // Output: Name: Alice
    }

    // Print the map (ordered by insertion)
    fmt.Println(om)  // Output: {"name": "Alice", "age": "30", "bio": "&lt;script&gt;alert(&#39;hi&#39;)&lt;/script&gt;"}

    // Disable HTML escaping
    om.SetEscapeHTML(false)
    fmt.Println(om)  // Output: {"name": "Alice", "age": "30", "bio": "<script>alert('hi')</script>"}

    // Delete a key
    om.Delete("age")
    fmt.Println(om)  // Output: {"name": "Alice", "bio": "<script>alert('hi')</script>"}

    // Clear all entries
    om.Clear()
    fmt.Println(om)  // Output: {}
}
```


## API

| Method               | Description                                  |
|----------------------|----------------------------------------------|
| `New(escapeHTML...)` | Create a new OrderedMap (escaping enabled by default) |
| `Set(key, value)`    | Add or update a key-value pair               |
| `Get(key)`           | Get value by key (returns value and existence flag) |
| `Delete(key)`        | Remove a key-value pair                      |
| `Clear()`            | Remove all key-value pairs                   |
| `SetEscapeHTML(on)`  | Enable/disable HTML escaping for string values |
| `String()`           | Get string representation (ordered by insertion) |


