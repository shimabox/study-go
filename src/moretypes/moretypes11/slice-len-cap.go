package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 長さを0に変更
	s = s[:0]
	printSlice(s)

	// 長さを4に拡張
	s = s[:4]
	printSlice(s)

	// 先頭2つの要素を削除。容量は4になる。
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
