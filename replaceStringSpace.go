package main

import (
	"fmt"
)

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))

}

func replaceSpace(s string) string {
	// 获取空格数量，string不能修改，但是可以读取每个字节
	spaceCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaceCnt++
		}
	}

	// 用新长度创建切片
	// 对于golang，数组长度是固定的，字符串是只读的。所以需要创建一个切片来存储新串，切片的长度还要在初始化时就确定，因为要从后向前填充元素。
	t := make([]byte, len(s)+spaceCnt*2)

	// i指向原串最后一位，j指向新切片最后一位。从后向前分别移动。
	// i==j，就是到"we"时
	for i, j := len(s)-1, len(t)-1; i <= j && i >= 0; i, j = i-1, j-1 {
		if s[i] != ' ' {
			t[j] = s[i]
		} else {
			t[j] = '0'
			t[j-1] = '2'
			t[j-2] = '%'
			j = j - 2 // 后面还有j--，所以j只需要-2
		}
	}

	return string(t)
}
