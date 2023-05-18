package main

import "fmt"

/**

707 设计链表，实现多种链表操作
https://mp.weixin.qq.com/s/Cf95Lc6brKL4g2j8YyF3Mg
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3

这道题目设计链表的五个接口：

获取链表第index个节点的数值
在链表的最前面插入一个节点
在链表的最后面插入一个节点
在链表第index个节点前面插入一个节点
删除链表的第index个节点
可以说这五个接口，已经覆盖了链表的常见操作，是练习链表操作非常好的一道题目
*/

var DummyHead listNode
var Size int

func main(){
	MyListNode()

	addAtTail(20)
	addAtHead(10)

	fmt.Println("-after add, List Node is:-")
	printListNode()

	fmt.Println("-get Node from ListNode:-")
	fmt.Println(getElementByIndex(0))
	fmt.Println(getElementByIndex(1))

	fmt.Println("-after add at index:-")
	addAtIndex(1, 0) // 头部之前添加
	addAtIndex(2, 1)  // 中间
	addAtIndex(30, 4) // 尾部后面添加
	addAtIndex(40, 6) // 超过元素数量的无效添加
	printListNode()

	fmt.Println("-after delete at index:-")
	deleteAtIndex(5) //  删除尾部元素之后不存在的数据
	deleteAtIndex(4)
	deleteAtIndex(1)
	deleteAtIndex(0)
	deleteAtIndex(-1)
	printListNode()

}

type listNode struct {
	Val int
	Next *listNode
}

func NewListNode() listNode{
	return listNode{
		Val:  0,
		Next: nil,
	}
}

func MyListNode(){
	DummyHead = NewListNode()
	Size = 0
}

// 头部增加Node（直接在DummyHead后追加）
func addAtHead(val int){
	newNode := NewListNode()
	newNode.Next = DummyHead.Next
	newNode.Val = val

	DummyHead.Next = &newNode
	Size++
}

// 尾部增加Node
func addAtTail(val int){
	// 如果对一个空链表在尾部插入，因为不存在头指针，所以不能通过i.Next获取到一个真实的链表元素，可能造成空指针错误。所以需要直接在虚拟头指针后直接插入一个node
	if DummyHead.Next == nil {
		addAtHead(val)
		return
	} else {
		newNode := NewListNode()
		newNode.Val = val

		// 存在真实头指针时，直接从真实头指针向后遍历，直到节点i来到最后一个真实节点(i.Next == nil)
		i := DummyHead.Next
		for ; i.Next != nil; i = i.Next {
			// 通过i = i.Next 以及 判断i.Next == nil跳出循环，得知当前i在最后一个元素上
		}

		i.Next = &newNode
		Size++
	}
}

// 特定位置增加Node(返回)
func addAtIndex(val int, index int) bool{
	if index > Size {
		//
		return false
	} else if index == Size {
		addAtTail(val)
	} else if index == 0 {
		addAtHead(val)
	} else {
		newNode := NewListNode()
		newNode.Val = val

		cur := &DummyHead
		for i:= 0; i < index; i++ {
			cur = cur.Next
		}

		newNode.Next = cur.Next
		cur.Next = &newNode
		Size++ // 其他过程有自己的Size运算
	}

	return true
}

func printListNode() {
	for i := DummyHead.Next; i != nil; i = i.Next{
		fmt.Println(i.Val)
	}
}

// 根据index(>=0)查询node
func getElementByIndex(index int) int{
	if index > Size - 1 || index < 0 {
		return -1
	}

	cur := &DummyHead
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func deleteAtIndex(index int){
	if index >= Size || index < 0 {
		return
	}

	if Size == 0 {
		return
	}

	if Size == 1 {
		DummyHead.Next = nil
		Size--
		return
	}

	cur := &DummyHead
	// 当index=0时，i<index-1的判定失败，不会进入循环，直接退出，cur还是虚拟头节点
	// 当index=1时，i = 0; i <= 0 成立，指针移动到index=0的位置（要删除的前一位）。下一轮循环判断i=1,i<=0失败退出，这一轮cur没变化。
	for i := 0; i <= index - 1; i++ {
		cur = cur.Next
	}

	// size >= 2 ，避免了cur.Next.Next出现空指针的情况
	cur.Next = cur.Next.Next
	Size--
}