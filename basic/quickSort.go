package main

import "fmt"

func main() {
    //arr := []int{47, 29, 71, 99, 78, 19, 24, 47}
    arr := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
    var ints = make([]int, len(arr))
    ints = arr[:]
    
    // pivot = ints[0]
    quick_sort(ints)
    fmt.Println(ints)
}

func quick_sort(ints []int) {
    i := 0
    l := len(ints)
    j := l - 1
    pivot_pointer := 0
    pivot := ints[pivot_pointer]
    for ;i < j; { // 循环=,for i<j
        for i < j && ints[j] > pivot {
            j = j - 1
        }
        if i < j {
            ints[j], ints[i] = ints[i], ints[j]
            i = i + 1
            pivot_pointer = j
        }

        for i < j && ints[i] < pivot {
            i = i + 1
        }
        if i < j {
            ints[i], ints[j] = ints[j], ints[i]
            j = j - 1
            pivot_pointer = i
        }
    }
    fmt.Println(ints)
    fmt.Println(pivot_pointer)
    if (pivot_pointer > 0) {
        quick_sort(ints[0:pivot_pointer])
    }
    if (pivot_pointer+1 < l) {
        quick_sort(ints[pivot_pointer+1:l])
    }
}
