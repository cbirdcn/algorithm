package main

// import "fmt"

// leetcode 950
// input: [2,13,3,11,5,17,7]
// output: [2,3,5,7,11,13,17]

func testCardsOrder() {
	//outputList := []int{1,2,3,4,5,6,7}
	// outputList2 := []int{1,2,3,4,5,6}
	//cardsOrder(outputList)
	// cardsOrder(outputList2)
}

//12
//12
//1.2(132)
//123
//1.2.(1324)
//1234
//1.2.3(15243:12-)
//12345
//1.2.3.(142635:123-465-56)
//123456
//1.2.3.4(1625374:123-4657-576-67)
//1234567
//1.2.3.4(15273648:1234-5768-687-78)
//12345678
//演算：1-2-3-4-,152-364-,1527364-,15273648
//192638475:1234-59687-6879-798-89:奇数个，中位数在最后一位，分隔开1234顺序挨个插入,6789依次在
//123456789x
//162x374958>12345-6x798-798x-8x9-9x
//演算过程:1-2-3-4-5-,162-374-5-,162-374958,162x374958

func cardsOrder(outputList []int) {
	// //inputList := make([]int, len(outputList))
	// tmpList := make([]int, len(outputList))
	// count := len(outputList)
	// //loopCount := 1
	// j := 0

	// 读取有序数组，一个个塞入对应的位置
	// for i:=0;i<count;i++ {
	// 	if 2 * i < count && tmpList[2 * i] == 0 {
	// 		tmpList[2 * i] = outputList[i]
	// 	}else {
	// 		for j=0;j<count;j++ {
	// 			if tmpList[j] == 0 {
	// 				tmpList[j] = outputList[i]
	// 				j += 
	// 			}
	// 		}
	// 	}
	// }

	//tmpList [2 * i] = outputList[i]

	//for i := 0; i < count; i++ {
	//	if 2 * i < count-1 {
	//		// 偶数个/奇数的前面数字
	//		tmpList[2 * i] = outputList[i]
	//	} else if 2 * i == count-1 {
	//		// 奇数个的最后一个数字
	//		tmpList[2 * i] = outputList[i]
	//	} else {
	//		loopCount++
	//		j = 2 + loopCount - 1
	//		//if count%2 == 0 {
	//		//	j = loopCount - 1
	//		//} else {
	//		//	j = 2 + loopCount - 1
	//		//}
	//		tmpList[j] = outputList[i]
	//		j = (j + loopCount * 2 + 1) % count
	//	}
	//}
	// fmt.Println(tmpList)

	//for i := 0; i<count; i++ {
	//	if 2 * i < count - 1 {
	//		// 偶数个/奇数的前面数字
	//		inputList[2 * i] = outputList[i]
	//	} else if 2 * i == count - 1 {
	//		// 奇数个的最后一个数字
	//		inputList[2 * i] = outputList[i]
	//	} else if i <= count - 1{
	//		loopCount++
	//		if count % 2 == 0 {
	//			i = loopCount -1
	//			continue
	//		} else {
	//			i = 2 + loopCount -1
	//			continue
	//		}
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println(inputList)
}



//type Node struct {
//	Value int
//	Next *Node
//}
//
//func NewNode() Node{
//	return Node{
//		Value: 0,
//		Next:  nil,
//	}
//}
//
//func cardsOrder(outputList []int){
//	//if len(outputList) <= 1 {
//	//	return outputList
//	//}else if len(outputList) == 2 {
//	//	sort.Ints(outputList)
//	//	return outputList
//	//} else if len(outputList) == 3 {
//	//	cursorCount := len(outputList) - 1
//	//	pre := cursorCount -1
//	//	next := cursorCount
//	//	inputList = append(inputList, outputList[0])
//	//	inputList = append(inputList, outputList[next])
//	//	inputList = append(inputList, outputList[pre])
//	//}
//
//	count := len(outputList)
//	// 创建slice，防止直接赋值触发panic，长度和原数组长度一致
//	//inputList = make([]int, count)
//
//	// 首轮取<=n/2的数字正序填入数组，如果长度是奇数，就填少于一半的奇数个数字，因为最后一个数字取出时会改变其他数字的顺序
//	firstCount := count / 2
//	//for i:=0;i<=firstCount;i++ {
//	//	inputList[i * 2] = outputList[i]
//	//}
//
//	// 其他数组逆向推导：让最后两个数字正序填入后，每向最前面加一个比这两个数字更小一点的数字（逆序添加），就要把最末尾的数字挪到加的数字之后，其他数字不变
//	listNode := Node{
//		Value: outputList[count-2],
//		Next:  nil,
//	}
//	listNode.Next = &Node{
//		Value: outputList[count-1],
//		Next:  nil,
//	}
//	// 存放头节点，最终要读取
//	headNode := NewNode()
//
//	for j:=count-1-firstCount-2;j>=0;j-- {
//		tmpNode := NewNode()
//		// 找到最后节点和倒数第二节点
//		for listNode.Next != nil {
//			if listNode.Next.Next == nil{
//				tmpNode = *listNode.Next
//				// 删除尾节点
//				listNode.Next = nil
//				break
//			}
//		}
//		// 尾结点作为新节点后接listNode
//		tmpNode.Next = &listNode
//		headNode = tmpNode
//		// 新建一个头节点，存放比原数字更小一点的数字
//		tmpNode = headNode
//		headNode.Value = outputList[j]
//		headNode.Next = &tmpNode
//	}
//
//	for i:=headNode.Next; i!=nil; i = i.Next {
//		fmt.Println(i.Value)
//	}
//
//}