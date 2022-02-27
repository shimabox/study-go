package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// sは[3 5 7]
	s = s[:2] // s = s[0:2] と同義
	fmt.Println(s)

	// sは[5 7]
	s = s[1:] // s = s[1:2] と同義
	fmt.Println(s)
}
