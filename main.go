package main

import "fmt"

func main() {

	type hoge struct {
		v int
	}

	x := hoge{
		v: 1,
	}

	y := hoge{
		v: 2,
	}

	fmt.Println(x)   //{1}
	fmt.Println(x.v) //1

	fmt.Println(&x)                      //&{1}
	fmt.Println("memory address:", &x.v) //0x140000180a0

	fmt.Println(y)   //{2}
	fmt.Println(y.v) //2

	fmt.Println(&y)                      //&{2}
	fmt.Println("memory address:", &y.v) //0x14000126010

	z := &hoge{
		v: 3,
	}

	fmt.Println(z)                       //&{3}
	fmt.Println(z.v)                     //3
	fmt.Println("memory address:", &z)   //0x14000120020
	fmt.Println("memory address:", &z.v) //0x14000126040

	fmt.Println(x.v)                     //1
	fmt.Println("memory address:", &x.v) //0x140000180a0
	fmt.Println(y.v)                     //2
	fmt.Println("memory address:", &y.v) //0x14000126010

	w := &hoge{
		v: 4,
	} //この段階ではまだ値が共有される環境にはなっていない

	fmt.Println(w)                       //&{4}
	fmt.Println(w.v)                     //4
	fmt.Println("memory address:", &w)   //0x14000120028
	fmt.Println("memory address:", &w.v) //0x14000126050

	fmt.Println(z)                       //&{3}
	fmt.Println(z.v)                     //3のまま
	fmt.Println("memory address:", &z)   //0x14000120020
	fmt.Println("memory address:", &z.v) //0x14000126040

	v := z //参照(ポインタ)が渡され、オブジェクトvとオブジェクトzの間で値が共有される環境になった

	fmt.Println(v.v) //3

	v.v = 100

	fmt.Println(z.v)                     //3ではなく100になる
	fmt.Println("memory address:", &z.v) //0x14000126040
	fmt.Println("memory address:", &v.v) //0x14000126040
	fmt.Println("memory address:", &z)   //0x14000120020
	fmt.Println("memory address:", &v)   //0x14000120030

	fmt.Println(w.v)                     //4
	fmt.Println("memory address:", &w.v) //0x14000126050

}
