/**
 * 以下の式で型変換を行うことができる
 *   <型名>(<変数>)
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// var x, y int = 3, 4
	x, y := 3, 4

	// f = √(x^2 + y^2) => 64ビット浮動小数点型に型変換
	var f float64 = math.Sqrt(float64(x*x + y*y))

	// z = uint型に型変換した f
	// なお var ステートメントで宣言した型と初期値の型が一致しない場合エラーとなる
	var z uint = uint(f)

	fmt.Println(x, y, f, z)
}
