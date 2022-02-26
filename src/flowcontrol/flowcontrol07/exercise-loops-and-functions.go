package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	tmp := z
	diff := 0.00000001 // こんなもんでいいのかな
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i+1, z)
		if math.Abs(tmp-z) < diff { // tmp-z < diff だと最初の1回目に左辺がマイナスになるので絶対値(math.Abs)で計算する
			break
		}
		tmp = z
	}
	return z
}

func main() {
	var x float64 = 6
	actual := Sqrt(x)
	expected := math.Sqrt(x)

	fmt.Println("Sqrt", actual)
	fmt.Println("math.Sqrt", expected)
	fmt.Println("result", math.Abs(expected-actual))
}
