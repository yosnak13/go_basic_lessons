package main

import "fmt"

//変数

var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	//明示的な定義
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	//複数の変数に代入
	var t, f bool = true, false
	fmt.Println(t, f)

	//まとめて定義
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	//変数だけ作りたい。初期値が入る
	var i3 int    //0
	var s3 string //空文字
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	//値の更新
	i = 150
	fmt.Println(i)

	//暗黙的な定義。関数内でしか暗黙的に定義できない
	i4 := 400
	fmt.Println(i4)
	//値の更新
	i4 = 450
	fmt.Println(i4)

	//i4 = "Hello"型推論、最初に定義した型のみ定義可能
	fmt.Println(i5)

	//関数呼び出し
	outer()
	//fmt.Println(s4) main関数内でs4は定義されていないため構文エラー

	//var s5 string = "Not used"
}
