package main

import "fmt"

//float(浮動小数点型)

func main() {
	var fl64 float64 = 2.4

	fmt.Println(fl64)

	//暗黙的定義の場合は自動でfloat64。intと違い環境依存ではない。
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	//float32は普段はあまり使われない。
	var fl32 float32 = 1.3
	fmt.Printf("%T\n", fl32)

	//型変換
	fmt.Printf("%T\n", float64(fl32))

	//プラスの無限大
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	//マイナスの無限大
	ninf := -1.0 / zero
	fmt.Println(ninf)

	//非数
	nan := zero / zero
	fmt.Println(nan)
}
