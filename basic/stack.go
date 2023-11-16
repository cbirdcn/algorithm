package main

import (
    "fmt"
    "sync"
    "errors"
    "fmt"
)

type T int // 包装类型，避免太多硬编码

type Stack struct {
    sync.RWMutex // 加锁操作
    array []T // 实际数据用slice存储
}

func NewStack() *Stack {
    stack := &Stack{} // 要返回指针，注意空struct及引用的写法
    stack.array = []T{} // 空struct，值为nil，一般用在返回前初始化
    return stack
}

func (s *Stack) Push(t T) { // 基于指针接收器的方法
    s.Lock() // 不要忘记加写锁
    s.array = append(s.array, t)
    s.UnLock() // 解锁
}

func (s *Stack) Pop() (t T, error){ // 需要返回出栈的值，以及错误
    // 需要判断是否为空，空要返回错误
    if s.IsEmpty() {
        return nil, errors.new("empty stack") // 
        // return nil, fmt.Errorf("empty stack") //
    }

    s.Lock()
    l := len(s.array)
    t = s.array(l-1)
    s.array = s.array(:l - 1)
    s.UnLock()
    return t, nil
}

func (s *Stack) Size() {
    s.RLock() // 读要加锁
    defer s.RUnlock() // 在return后defer
    return len(s.array)
}

func (s *Stack) IsEmpty() {
    return s.Size() == 0
}

func main() {
    s := NewStack()
    s.Push(1)
    s.Push(2)
    fmt.Println(s.Empty())
    fmt.Println(s.Size())
    fmt.Println(s.array())
    s.Pop()
    fmt.Println(s.array())
}
