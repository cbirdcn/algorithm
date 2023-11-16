package main

import "fmt"

func main() {
    ints := []int{49,38,65,97,76,13,27,49}
    
    l := len(ints)
    for i := 0; i < l; i = i + 1 {
        for j := 0; j < l - i -1; j++ {
            if ints[j] > ints[j+1] {
		    ints[j], ints[j+1] = ints[j+1], ints[j]
            }
        }
    }
    fmt.Println(ints)
}
