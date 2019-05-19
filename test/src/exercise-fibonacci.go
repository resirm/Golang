package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	l, ll := 0, 1
	return func() int {
		ret := l
		l, ll = ll, l+ll
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
