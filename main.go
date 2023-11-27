package main

import "fmt"

type Hoge struct {
	value int
}

// ポインタ型ではなく値型なので、h.valueを書き換えられない
func (h Hoge) HogeA(v int) {
	h.value = v + 5
}

// ポインタ型なので、h.valueを書き換えられる
func (h *Hoge) HogeB(v int) {
	h.value = v + 5
}

func main() {
	var h Hoge

	h.HogeA(0)
	// 0(ゼロ値)が出力される
	fmt.Printf("HogeA(ポインタ型ではない)の出力結果:%v\n", h.value)

	h.HogeB(0)
	// +5された値が出力される
	fmt.Printf("HogeB(ポインタ型である)の出力結果:%v\n", h.value)
}
