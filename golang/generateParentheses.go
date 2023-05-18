package main

import "fmt"

// 输入n，生成n对左右有序的括号对组合
// 如输入3，生成((()))、(()())、(())()、()(())、()()()


var res []string

func main () {
	n := 3
	path := ""
	res = make([]string, 0)

	backTrack(n, 0, 0, path)
	fmt.Println(res)
}

func backTrack(n int, l int, r int, path string) {
	if l == n && r == n {
		res = append(res, path)
	} else {
		if l < n {
			backTrack(n, l + 1, r, path+"(")
		}
		if r < l {
			backTrack(n, l, r + 1, path+")")
		}
	}
}
