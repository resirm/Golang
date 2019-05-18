package main

import (
	"fmt"
)

type Mp map[int]string

func test(m Mp) {
	fmt.Printf("test: %T, %v\n", m, m)
	m[1] = "changed"
	fmt.Printf("test: %T, %v\n", m, m)
}

func main() {
	var mp = make(Mp)
	fmt.Printf("main: %T, %v\n", mp, mp)
	test(mp)
	fmt.Printf("main: %T, %v\n", mp, mp)

}
