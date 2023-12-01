package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	fmt.Printf("最初の変数s:%v, 最初の変数sのアドレス:%v\n", s, &s[1]) //最初の変数s:[1 2 3] 最初の変数sのアドレス:0x1400012c020
	updateArgSlice(s)
	fmt.Printf("サブ関数処理後の変数s:%v\n", s) //サブ関数処理後の変数s:[1 2 3]
}

func updateArgSlice(s []int) {
	fmt.Printf("要素追加前のアドレス:%v\n", &s[1]) //素追加前のアドレス:0x1400012c020

	//要素を追加した時点で配列のサイズが変わり、「配列のサイズ」という型も変わる
	//よってこの配列を参照するスライスのアドレスも変わるので、メイン関数内には影響を及ぼせない
	s = append(s, 4)
	fmt.Printf("要素追加後のアドレス:%v\n", &s[1]) //要素追加後のアドレス:0x14000132008
	s[0] = 10
}
