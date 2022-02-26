package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // iへのポインタを引き出す
	fmt.Println(*p) // ポインタpを通してiから値を読みだす
	*p = 21         // ポインタpを通してiへ値を代入する
	fmt.Println(i)  // iの新しい値を参照

	p = &j          // jへのポインタを引き出す
	fmt.Println(p)  // メモリアドレス
	*p = *p / 37    // jをポインタで割る
	fmt.Println(j)  // jの新しい値を参照
	fmt.Println(*p) // * オペレータは、ポインタの指す先の変数を示す
}
