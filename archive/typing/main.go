package main

import "fmt"

//string

func main() {
	var s string = "Hello Go"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	//複数行文字列はバッククォーテーションを使用
	fmt.Println(`test
	test
		test`)
	//文字列としてのダブルクォーテーションは２通り
	fmt.Println("\"")
	fmt.Println(`""`)

	//文字列の1文字目を取得。
	//byte配列になっていて、Hは72
	//stringを指定して、その引数に取得したい対象を指定すると文字列になる。
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}
