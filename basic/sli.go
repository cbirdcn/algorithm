package main

import "fmt"
import "slices"

func main() {
    s0 := []int{1, 2, 3} // 切片的简短声明赋值，和数组类似，但是没有长度。但是一般不这样用
    fmt.Println("emp:", s0, "len:", len(s0), "cap:", cap(s0))
    
    var s []int
    s = make([]int, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    
    s = append(s, 1)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    
    
    var s2 []int
    s2 = make([]int, 0, 3) // 初始化cap=3，但是不初始化，所以len=0
    s2 = append(s2, 1, 2)
    fmt.Println("emp:", s2, "len:", len(s2), "cap:", cap(s2)) // 追加2个元素后，len=2
    // emp: [1 2] len: 2 cap: 3

    s = append(s, s2...)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    
    var c []int
    c = make([]int, len(s))
    fmt.Println("before cpy:", c, "len:", len(c), "cap:", cap(c))
    copy(c, s) // 没有返回值
    fmt.Println("cpy1:", c, "len:", len(c), "cap:", cap(c))
        var c2 []int
    c2 = make([]int, len(s)-1, len(s)) // 
    copy(c2, s)
    fmt.Println("cpy2:", c2, "len:", len(c2), "cap:", cap(c2))
    var c3 []int
    c3 = make([]int, len(s)-1)
    copy(c3, s) // 没有返回值
    fmt.Println("cpy3:", c3, "len:", len(c3), "cap:", cap(c3))
    
    low := 0
    high := len(s)
    var p []int
    p = s[low:high]
    fmt.Println("part:", p, "len:", len(p), "cap:", cap(p))
    // part: [0 0 0 1 1 2] len: 6 cap: 6
    p = s[:high]
    fmt.Println("part:", p, "len:", len(p), "cap:", cap(p))
    // part: [0 0 0 1 1 2] len: 6 cap: 6
    p = s[:]
    fmt.Println("part:", p, "len:", len(p), "cap:", cap(p))
    
    if slices.Equal(s, p) {
        fmt.Println("s == p")
    }
    // part: [0 0 0 1 1 2] len: 6 cap: 6
    // p = s[:-1]
    // fmt.Println("part:", p, "len:", len(p), "cap:", cap(p))
    // part: [0 0 0 1 1 2] len: 6 cap: 6
    // part: [0 0 0 1 1 2] len: 6 cap: 6
    //     var c4 []int
    // c4 = make([]int, cap(s)+1) // 拷贝到更大的dst中
    // copy(c4, s)
    // fmt.Println("cpy4:", c4, "len:", len(c4), "cap:", cap(c4))
    //     var s3 []string
    // s3 = make([]string, 0, 3)
    // s3 = append(s3, "a", "b", "c")
    // fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
    // var c []string
    // c = make([]string, len(s3))
    // copy(c, s3)
    // fmt.Println("cpy:", c, "len:", len(c), "cap:", cap(c))
    
    // var a [5]int // 注意数组类型包含长度（不带长度是切片slice），写到变量后，长度在前元素类型在后。遇到特殊类型比如二维数组就是[m][n]int
    // fmt.Println("emp:", a) // 默认为零值，int类型意味着都是0
    // fmt.Println("emp:", a[0])
    // // fmt.Println("emp:", a[5]) // 越界
    // // 输出
    // // emp: [0 0 0 0 0]
    // // emp: 0
    // // invalid array index 5 (out of bounds for 5-element array)

    // a = [5]int{1, 2, 3, 4, 5} // 初始化用{}，类型一致
    
    // change(&a)
    // fmt.Println(a)
    
    
    // fmt.Println("Hello world!")
    // a := 1
    // b := 2
    
    
    
    // replaceFalse(a, b)
    // fmt.Println(a, b)
    
    // fmt.Println(&a)
    // replaceTrueInC(&a, &b)
    // fmt.Println(&a)
    // fmt.Println(a, b)
    
    //     fmt.Println(&a)
    // replaceTrueInGo(&a, &b)
    // fmt.Println(&a)
    // fmt.Println(a, b)
}

func change(x *[5]int) {
    x[0] = 100
    // *x[1] = 200 // invalid indirect of x[1] (type int)
    (*x)[2] = 300
}

func replaceFalse(x, y int) {
    x, y = y, x
}

func replaceTrueInC(x, y *int) {
    var tmp int
    tmp = *x
    *x = *y
    *y = tmp
}

func replaceTrueInGo(x, y *int) {
    *x, *y = *y, *x
}
