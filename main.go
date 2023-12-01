package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Printf("address is %v\n", &i) //0x14000126008
	updateArgInt(i)                   //この段階の実引数iは上記の変数iと同じアドレス？
	fmt.Printf("引数がint型の時、サブ関数の処理後、iの値は%v\n", i)
	updateArgPointer(&i)
	fmt.Printf("引数が*int型の時、サブ関数の処理後、iの値は%v\n", i)
}

// 引数iのアドレス：0x14000126020
// 値渡しなのでサブ関数のスコープ内でしかインクリメントされない
func updateArgInt(i int) {
	fmt.Printf("address is %v\n", &i)
	i++
}

// 引数iに渡ったアドレス：0x14000126008
// 参照渡しなのでサブ関数を越えて影響する
func updateArgPointer(i *int) {
	fmt.Printf("address is %v\n", i)
	*i++
}
