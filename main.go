package main

import "fmt"

type Food struct {
	ID          int
	Name        string
	Description string
}

func main() {
	fmt.Println(MakeResult())           //map[うまいもの:0x14000106270]
	fmt.Println(MakeResult()["うまいもの"])  //&{1 焼肉 すこぶるうまい}
	fmt.Println(*MakeResult()["うまいもの"]) //{1 焼肉 すこぶるうまい}

	x := MakeResult()
	x["うまいもの"].Description = "hoge"
	fmt.Println(x)                      //map[うまいもの:0x14000112300]
	fmt.Println(x["うまいもの"].ID)          //1
	fmt.Println(x["うまいもの"].Name)        //焼肉
	fmt.Println(x["うまいもの"].Description) //hoge

	y := MakeResult1()
	//エラー：cannot assign to struct field y["うまいもの"].Description in map
	// y["うまいもの"].Description = "fuga"

	fmt.Println(y) //map[うまいもの:{1 焼肉 すこぶるうまい}]
}

// エラー：cannot assign to struct field result["うまいもの"].Description in map
/*
func MakeResult() map[string]Food {
	result := make(map[string]Food)
	result["うまいもの"] = Food{
		ID:   1,
		Name: "焼肉",
	}
	result["うまいもの"].Description = "すこぶるうまい"

	return result
}
*/

func MakeResult() map[string]*Food {
	result := make(map[string]*Food)
	result["うまいもの"] = &Food{
		ID:   1,
		Name: "焼肉",
	}
	result["うまいもの"].Description = "すこぶるうまい"

	return result
}

func MakeResult1() map[string]Food {
	pointerResult := make(map[string]*Food)
	pointerResult["うまいもの"] = &Food{
		ID:   1,
		Name: "焼肉",
	}
	pointerResult["うまいもの"].Description = "すこぶるうまい"

	result := make(map[string]Food)
	for k, v := range pointerResult {
		result[k] = *v
	}
	return result
}
