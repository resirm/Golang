package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walkr(t, ch)
	close(ch)
}

func Walkr(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walkr(t.Left, ch)
	ch <- t.Value
	Walkr(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var ar1, ar2 []int
	for c := range ch1 {
		fmt.Println(c)
		ar1 = append(ar1, c)
	}
	for c := range ch2 {
		fmt.Println(c)
		ar2 = append(ar2, c)
	}
	for idx := range ar1 {
		if ar1[idx] != ar2[idx] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// Reference code
// package main

// import "fmt"
// import "code.google.com/p/go-tour/tree"

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
//     var walker func(t *tree.Tree)
//     walker = func (t *tree.Tree) {
//         if (t == nil) {
//             return
//         }
//         walker(t.Left)
//         ch <- t.Value
//         walker(t.Right)
//     }
//     walker(t)
//     close(ch)
// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
//     ch1, ch2 := make(chan int), make(chan int)

//     go Walk(t1, ch1)
//     go Walk(t2, ch2)

//     for {
//         v1,ok1 := <- ch1
//         v2,ok2 := <- ch2

//         if v1 != v2 || ok1 != ok2 {
//             return false
//         }

//         if !ok1 {
//             break
//         }
//     }

//     return true
// }

// func main() {
//     fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
//     fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))

// }
