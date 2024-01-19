package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["k1"] = 1
    m["k2"] = 2
    fmt.Println(m)
    fmt.Println(len(m))
    delete(m, "k2")

    v1, ok := m["k1"]
    fmt.Println(v1, ok)
    // 1 true
    v2, ok := m["k2"]
    fmt.Println(v2, ok)
    // 0 false
    
    clear(m)
    fmt.Println(m)
}
