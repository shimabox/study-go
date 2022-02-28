package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// 最初の2つの数は 0, 1
	a1, a2 := 0, 1
	return func() (ans int) {
		ans = a1
		// 3つ目以降は前の2つの数の和を並べたもの
		a1, a2 = a2, a1+a2
		// fmt.Println(ans, a1, a2, a1+a2)
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
