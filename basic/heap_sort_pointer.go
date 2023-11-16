package main

import "fmt"

func main() {
    arr := [...]int{20,30,90,40,70,110,60,10,100,50,80}
    // arr := [...]int{9,8,7,6,5,4,3,2,1}
    var ints []int = make([]int, 0, len(arr))
    ints = arr[:]

    heap_sort(&ints)

    fmt.Println(ints)
}

func heap_sort(ints *[]int) {
    for k := len(*ints); k >= 2; k-- {
        // 构建大顶堆：从第一个非叶子结点(n/2-1)从下至上，从右至左（i--）调整结构。for结束会在ints[0]得到一个最大的稳定节点
        for i := k / 2 - 1; i >= 0; i = i - 1 {
            heap_adjust(ints, i, k - 1)
        }
        // 每次大顶堆构建结束，从第n个节点到第2个都要做首尾交换操作，然后对剩下的节点再次构建大顶堆
        (*ints)[k - 1], (*ints)[0] = (*ints)[0], (*ints)[k - 1]
    }
}

func heap_adjust(ints *[]int, i, k int) {
    // 2i+1是左子节点
    if (2 * i + 1) <= k && (*ints)[2 * i + 1] > (*ints)[i] {
        (*ints)[i], (*ints)[2 * i + 1] = (*ints)[2 * i + 1], (*ints)[i]
    }
    // 2i+2是右子节点
    if (2 * i + 2) <= k && (*ints)[2 * i + 2] > (*ints)[i] {
        (*ints)[i], (*ints)[2 * i + 2] = (*ints)[2 * i + 2], (*ints)[i]
    }
}
