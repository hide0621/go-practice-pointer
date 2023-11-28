package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4}
	s2 := s1

	fmt.Println(&s1) //&[1 2 3 4]
	fmt.Println(&s2) //&[1 2 3 4]

	fmt.Println(s1)     // [1 2 3 4]
	fmt.Println(s2)     // [1 2 3 4]
	fmt.Println(&s1[0]) // 0xc000100000
	fmt.Println(&s2[0]) // 0xc000100000

	s2[2] = 5

	fmt.Println(s1)     // [1 2 5 4]
	fmt.Println(s2)     // [1 2 5 4]
	fmt.Println(&s1[2]) // 0xc000100010
	fmt.Println(&s2[2]) // 0xc000100010

}
