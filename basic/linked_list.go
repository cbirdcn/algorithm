package main

import (
    "fmt"
    "sync"
    "errors"
)
type T int

type Node struct {
    val T
    next *Node
}

type LinkedList struct {
    sync.RWMutex
    head *Node
    size int
}

func NewNode(t T) (*Node){ // 返回节点的指针
    return &Node{ val: t} // 需要提供节点的值，不知道next指向所以不提供
}

func NewLinkedList() (*LinkedList){ // 返回指针
    return &LinkedList{} // 初始化空链表，不需要首个节点情况
}

func (l *LinkedList) Add(t T) {
    l.Lock()
    defer l.Unlock()

    node := NewNode(t)
    
    // 需要区分是否头节点
    if l.size == 0 {
        l.head = node
    } else {
        head := l.head
        for head.next != nil {
            head = head.next
        }
        head.next = node
    }

    // 长度+1
    l.size = l.size + 1
}

func (l *LindedList) Traverse() {
    l.RLock()
    defer l.RUnlock()

    head := l.head
    for head.next != nil {
        head.Print() // 打印节点
        head = head.next
    }
    head.Print() // 打印最后一个节点
}

func (n *Node) Print() {
    fmt.Println(n.val)
}

func main() {
    l := NewLinkedList()
    l.Add(1)
    l.Traverse()
}
