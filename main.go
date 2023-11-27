package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func NewPersonValue(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	//ポインタ型のインスタンスを作成
	personPointer := NewPersonPointer("山田", "太郎", 30)

	fmt.Println("First Name (Pointer):", personPointer.FirstName)
	fmt.Println("Last Name (Pointer):", personPointer.LastName)
	fmt.Println("Age (Pointer):", personPointer.Age)

	//値型のインスタンスを作成
	personValue := NewPersonValue("佐藤", "二郎", 25)

	fmt.Println("\nFirst Name (Value):", personValue.FirstName)
	fmt.Println("Last Name (Value):", personValue.LastName)
	fmt.Println("Age (Value):", personValue.Age)
}
