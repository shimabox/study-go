package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("called") // 呼び出されたときに1回呼ばれる
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
