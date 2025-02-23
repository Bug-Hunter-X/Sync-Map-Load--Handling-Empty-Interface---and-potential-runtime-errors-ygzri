```go
package main

import (

    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    m.Store("key", "value")
    value, ok := m.Load("key")
    if !ok {
        fmt.Println("Key not found")
    } else {
        strValue, ok := value.(string)
        if !ok {
            fmt.Println("Type assertion failed")
        } else {
            fmt.Println(strValue)
        }
    }
}
```