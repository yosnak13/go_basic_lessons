package main

import "fmt"

//int型

func main() {
	var i int = 100
	/*
		最大値 最小値
		int8
		int16
		int32
		int64
	*/

	var i2 int64 = 200

	fmt.Println(i + 50)
	// 以下は型違いのためエラーになる
	//fmt.Println(i + i2)

	//書式指定子（型を表示）
	fmt.Printf("%T\n", i2)

	//型変換
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))
}
Footer

