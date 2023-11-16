package main

import "fmt"

func main() {
    arr1 := [...]int{4,5,7,8}
    var ints1 = make([]int, 4)
    ints1 = arr1[:]
    arr2 := [...]int{1,2,3,6}
    var ints2 = make([]int, 4)
    ints2 = arr2[:]

    ints := make([]int, 0, len(arr1) + len(arr2))
    mergeSort(ints1, ints2, &ints)
    fmt.Println(ints)
}

func mergeSort(ints1, ints2 []int, ints_pointer *[]int) {
    l1 := len(ints1)
    l2 := len(ints2)
    i:=0
    j:=0
    for i < l1 && j < l2 {
        if ints1[i] <= ints2[j] {
            *ints_pointer = append(*ints_pointer, ints1[i])
            i = i + 1
            continue
        } else {
            *ints_pointer = append(*ints_pointer, ints2[j])
            j = j + 1
            continue
        }
    }
    if i == l1 {
        *ints_pointer = append(*ints_pointer, ints2[j:l2]...)
    }
    if j == l2 {
        *ints_pointer = append(*ints_pointer, ints1[i:l1]...)
    }
}
