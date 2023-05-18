package main

/*
通过container包的list(链表)，实现队列、栈
https://blog.csdn.net/qq_43684922/article/details/116381764
*/

import (
	"container/list"
	"fmt"
)

func main() {
	// 1.队列
	// 生成队列
	l := list.New()
	// 入队
	l.PushBack(1)
	l.PushBack(2)
	// 出队
	i1 := l.Front()
	l.Remove(i1)
	fmt.Printf("%d\n", i1.Value)
	i2 := l.Front()
	l.Remove(i2)
	fmt.Printf("%d\n", i2.Value)

	// 2.栈
	//入栈
	l.PushBack(1)
	l.PushBack(2)
	// 出栈
	i3 := l.Back()
	l.Remove(i3)
	fmt.Printf("%d\n", i3.Value)
	i4 := l.Back()
	l.Remove(i4)
	fmt.Printf("%d\n", i1.Value)
}

/*
运行:
1
2
2
1
*/
