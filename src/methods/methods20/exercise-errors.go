package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) { // methods20/exercise-loops-and-functions.go
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := float64(1)
	tmp := z
	diff := 0.00000001
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i+1, z)
		if math.Abs(tmp-z) < diff {
			break
		}
		tmp = z
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	// fmtパッケージは、変数を文字列で出力する際にerrorインタフェースを確認するので変換する
	v := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt nagative number: %v", v)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
