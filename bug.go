```go
func main() {
    var m sync.Map
    m.Store("key", "value")
    value, ok := m.Load("key")
    if !ok {
        fmt.Println("Key not found")
    }
    fmt.Println(value)
}
```